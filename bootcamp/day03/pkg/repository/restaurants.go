package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"

	"server/types"
	"server/types/models"
	"server/utils"
)

type RestaurantsRepository struct {
	db      *gorm.DB
	elastic *elasticsearch.Client
}

func NewRestaurantsRepository(
	db *gorm.DB,
	elasic *elasticsearch.Client,
) *RestaurantsRepository {
	return &RestaurantsRepository{
		db:      db,
		elastic: elasic,
	}
}

func (r *RestaurantsRepository) Create(restaurant models.Restaurant) (string, error) {
	res := r.db.Create(&restaurant)
	if res.Error != nil {
		return "0", res.Error
	}
	restaurantBytes, err := json.Marshal(restaurant)
	if err != nil {
		return "0", err
	}
	elasticRes, err := r.elastic.Index(
		types.GLOABAL_CONFIG.RestaurantsIndex,
		bytes.NewReader(restaurantBytes),
		r.elastic.Index.WithDocumentID(restaurant.ID),
	)
	if err != nil || elasticRes.IsError() {
		return "0", err
	}
	return restaurant.ID, nil
}

func (r *RestaurantsRepository) Delete(id string) (bool, error) {
	res := r.db.Delete(&models.Restaurant{ID: id})
	if res.Error != nil {
		return false, res.Error
	}
	elasticRes, err := r.elastic.Delete(types.GLOABAL_CONFIG.RestaurantsIndex, id)
	if err != nil {
		return false, err
	}
	if elasticRes.IsError() {
		return false, errors.New(elasticRes.Status())
	}

	return true, nil
}

func (r *RestaurantsRepository) GetAll(
	pagination types.PagingIncoming,
) (types.PaginatedRestaurants, error) {
	restaurantsRes := types.PaginatedRestaurants{}
	restaurantsRes.Paging = types.PagingOutgoing{PagingIncoming: pagination, Total: 0}

	total := int64(0)
	totalRes := r.db.Model(&models.Restaurant{}).Count(&total)

	if totalRes.Error != nil {
		log.Fatalf(
			"failed to read restaurants with paging: %v\n",
			totalRes.Error.Error(),
		)
		return types.PaginatedRestaurants{}, totalRes.Error
	}
	restaurantsRes.Paging.Total = total

	pageCount := int(
		math.Ceil(float64(restaurantsRes.Paging.Total) / float64(restaurantsRes.Paging.PageSize)),
	)
	if restaurantsRes.Paging.PageNumber > pageCount {
		return restaurantsRes, errors.New(
			fmt.Sprintf(
				"Only have %d pages, requested %d",
				pageCount,
				restaurantsRes.Paging.PageNumber,
			),
		)
	}

	res := r.db.Scopes(utils.Paginate(pagination)).Preload("Location").Find(&restaurantsRes.Items)

	if res.Error != nil {
		log.Fatalf(
			"failed to read restaurants with paging: %v, error: %v\n",
			pagination,
			res.Error.Error(),
		)
		return restaurantsRes, res.Error
	}

	return restaurantsRes, nil
}

func (r *RestaurantsRepository) GetById(id string) (models.Restaurant, error) {
	res := models.Restaurant{}
	found := int64(0)
	res.ID = id
	quey := r.db.Preload("Location").Find(&res)
	quey.Count(&found)
	if found == 0 {
		return res, errors.New(fmt.Sprintf("Restaurant with id %s not found", id))
	}

	if quey.Error != nil {
		return res, quey.Error
	}
	return res, nil
}

func (r *RestaurantsRepository) Search(
	query string,
	paging types.PagingIncoming,
) (types.PaginatedRestaurants, error) {
	indexName := types.GLOABAL_CONFIG.RestaurantsIndex // TODO use config
	pagingRest := types.PagingOutgoing{}
	queryString := fmt.Sprintf(`
{
    "from": %d,
    "size": %d,
  "query": {
    "multi_match" : {
      "query":    "%s", 
      "fields": [ "name", "address", "phone" ] 
    }
  }
}
  `, (paging.PageNumber-1)*paging.PageSize, paging.PageSize, query)

	res, err := r.elastic.Search(
		r.elastic.Search.WithContext(context.Background()),
		r.elastic.Search.WithIndex(indexName),
		r.elastic.Search.WithBody(strings.NewReader(queryString)),
		r.elastic.Search.WithPretty(),
	)
	if err != nil {
		return types.PaginatedRestaurants{}, err
	}

	defer res.Body.Close()
	var hits struct {
		Hits struct {
			Hits []struct {
				Source models.Restaurant `json:"_source"`
			} `json:"hits"`
			Total struct {
				Value int64 `json:"value"`
			} `json:"total"`
		} `json:"hits"`
	}
	err = json.NewDecoder(res.Body).Decode(&hits)
	if err != nil {
		return types.PaginatedRestaurants{}, err
	}
	resSlice := []models.Restaurant{}
	for _, hit := range hits.Hits.Hits {
		resSlice = append(resSlice, hit.Source)
	}
	pagingRest.PageNumber = paging.PageNumber
	pagingRest.Total = hits.Hits.Total.Value
	pagingRest.PageSize = paging.PageSize
	pageCount := int(math.Ceil(float64(pagingRest.Total) / float64(pagingRest.PageSize)))
	if pagingRest.PageNumber > pageCount {
		return types.PaginatedRestaurants{}, errors.New(
			fmt.Sprintf(
				"page number out of bounds. Only have %d pages.",
				pageCount,
			),
		)
	}

	return types.PaginatedRestaurants{
		Paging: pagingRest,
		Items:  resSlice,
	}, nil
}

func (r *RestaurantsRepository) Update(restaurant models.Restaurant) (string, error) {
	res := r.db.Save(&restaurant)
	if res.Error != nil {
		return "0", res.Error
	}
	restaurantJsonBytes, err := json.Marshal(struct {
		Doc models.Restaurant `json:"doc"`
	}{
		Doc: restaurant,
	})
	if err != nil {
		return "0", err
	}

	log.Printf("resid: %s, index: %s", restaurant.ID, types.GLOABAL_CONFIG.RestaurantsIndex)
	elasticRes, err := r.elastic.Update(
		types.GLOABAL_CONFIG.RestaurantsIndex,
		restaurant.ID,
		bytes.NewReader(restaurantJsonBytes),
		r.elastic.Update.WithErrorTrace(),
	)
	if err != nil || elasticRes.IsError() {
		errs := make([]byte, 10000)
		_, err := elasticRes.Body.Read(errs)
		if err != nil {
			log.Println(err.Error())
		}
		elasticRes.Body.Close()
		return "0", err
	}

	return restaurant.ID, nil
}

func (r *RestaurantsRepository) GetClosest(
	target models.Restaurant,
	paging types.PagingIncoming,
	distance int,
) (types.PaginatedRestaurants, error) {
	restaurants := types.PaginatedRestaurants{}
	pagingRest := types.PagingOutgoing{}
	query := utils.GetFindClosestRestaurantQuery(
		paging,
		target.Location.Lat,
		target.Location.Lon,
		distance,
	)
	println("QUERY::::::: ")
	println(query)
	res, err := r.elastic.Search(
		r.elastic.Search.WithContext(context.Background()),
		r.elastic.Search.WithIndex(),
		r.elastic.Search.WithBody(strings.NewReader(query)),
		r.elastic.Search.WithPretty(),
	)
	if err != nil {
		return restaurants, err
	}

	defer res.Body.Close()
	var hits struct {
		Hits struct {
			Hits []struct {
				Source models.Restaurant `json:"_source"`
			} `json:"hits"`
			Total struct {
				Value int64 `json:"value"`
			} `json:"total"`
		} `json:"hits"`
	}
	err = json.NewDecoder(res.Body).Decode(&hits)
	if err != nil {
		return restaurants, err
	}

	for _, hit := range hits.Hits.Hits {
		restaurants.Items = append(restaurants.Items, hit.Source)
	}

	pagingRest.PageNumber = paging.PageNumber
	pagingRest.Total = hits.Hits.Total.Value
	pagingRest.PageSize = paging.PageSize
	pageCount := int(math.Ceil(float64(pagingRest.Total) / float64(pagingRest.PageSize)))
	if pagingRest.PageNumber > pageCount {
		return restaurants, errors.New(
			fmt.Sprintf("found only %d pages. Asked for %d", pageCount, paging.PageNumber),
		)
	}
	restaurants.Paging = pagingRest
	return restaurants, nil
}
