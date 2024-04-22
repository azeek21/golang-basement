package restaurants

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"

	"server/core/paging"
)

func QueryRestaurants(
	es *elasticsearch.Client,
	indexName string,
	text string,
	page paging.PagingIncoming,
) ([]Restaurant, paging.PagingOutgoing, error) {
	pagingRest := paging.PagingOutgoing{}
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
  `, page.PageNumber*page.PageSize, page.PageSize, text)

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(indexName),
		es.Search.WithBody(strings.NewReader(queryString)),
		es.Search.WithPretty(),
	)
	if err != nil {
		return nil, pagingRest, err
	}

	defer res.Body.Close()
	var hits struct {
		Hits struct {
			Hits []struct {
				Source Restaurant `json:"_source"`
			} `json:"hits"`
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
		} `json:"hits"`
	}
	err = json.NewDecoder(res.Body).Decode(&hits)

	if err != nil {
		return nil, pagingRest, err
	}
	resSlice := []Restaurant{}
	for _, hit := range hits.Hits.Hits {
		resSlice = append(resSlice, hit.Source)
	}
	pagingRest.PageNumber = page.PageNumber
	pagingRest.Total = hits.Hits.Total.Value
	pagingRest.PageSize = page.PageSize

	return resSlice, pagingRest, nil
}

func RegisterRestaurantControllers(enginge *gin.Engine, es *elasticsearch.Client) error {
	group := enginge.Group("/")
	group.GET("/", func(c *gin.Context) {
		pagination, err := paging.GetPaginationFromQueryParams(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		query := c.DefaultQuery("q", "")
		restaurants, resPaging, err := QueryRestaurants(es, "places", query, pagination)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.HTML(200, "index.html", gin.H{"places": gin.H{
			"query": query,
			"items": restaurants,
			"paging": gin.H{
				"next":                 resPaging.PageNumber + 1,
				"prev":                 resPaging.PageNumber - 1,
				"total":                resPaging.Total,
				paging.PAGE_NUMBER_KEY: resPaging.PageNumber,
			},
		}})
	})
	return nil
}
