package initializers

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"server/types"
	"server/types/models"
)

var POSTGRES_LIMIT = 5000

var RestaurantMappings = `
{
  "settings": {
    "number_of_shards": 1,
    "max_result_window": 20000
  },
  "mappings": {
    "properties": {
      "id": {
        "type": "text"
      },
      "name": {
        "type": "text"
      },
      "id": {
        "type": "text"
      },
      "phone": {
        "type": "text"
      },
      "address": {
        "type": "text"
      },
      "location": {
        "type": "geo_point"
      }
    }
  }
}
`

func MigrateRestaurantsElastic(
	elastic *elasticsearch.Client,
	config *types.RestaurantsConfig,
	restaurants []models.Restaurant,
) error {
	err := CreateRestaurantsIndex(elastic, config)
	if err != nil {
		return err
	}

	err = InitRestaurantsElastic(elastic, restaurants, *config)
	if err != nil {
		return err
	}
	log.Println("INDEXED places in elasticsearch")

	return nil
}

func MigrateRestaurantsPostgres(db *gorm.DB, restaurants []models.Restaurant) error {
	err := db.Migrator().DropTable(&models.User{}, &models.Location{}, &models.Restaurant{})
	if err != nil {
		return err
	}
	log.Println("DROP TABLE restaurants, locations, users")

	if err := db.AutoMigrate(&models.User{}, &models.Restaurant{}, &models.Location{}); err != nil {
		return err
	}
	log.Println("CREATED TABLES/SCHEMAS restaurants, locations; users;")

	start, end := 0, len(restaurants)
	for {
		tmpEnd := start + POSTGRES_LIMIT
		if tmpEnd > end {
			tmpEnd = end
		}
		chunk := restaurants[start:tmpEnd]
		start = tmpEnd
		res := db.Session(&gorm.Session{SkipHooks: true}).Create(chunk)
		log.Printf("[postgres]: inserted %v out of %v\n", tmpEnd, end)
		if res.Error != nil {
			return res.Error
		}
		if start >= end {
			break
		}
	}
	log.Println("MIGRATED RESTAURANT MODELS: ", "restaurants", "locations")
	return nil
}

func CreateRestaurantsIndex(client *elasticsearch.Client, config *types.RestaurantsConfig) error {
	res, err := client.Indices.Delete([]string{config.IndexName})

	if err != nil || res.IsError() {
		if res.StatusCode != 404 {
			return errors.New(
				fmt.Sprintf(
					"error [deleting old index: %s]: %v\n",
					config.IndexName,
					res.String(),
				),
			)
		} else {
			return err
		}
	}
	log.Println("DELETED places index from elasticsearch")

	res, err = client.Indices.Create(
		config.IndexName,
		client.Indices.Create.WithBody(strings.NewReader(RestaurantMappings)),
	)

	if res.IsError() {
		return errors.New(fmt.Sprintf("error [creating restaurants index]: %v\n", res.Body))
	}

	if err != nil {
		return err
	}

	log.Println("CREATED places index in elasticsearch")
	return nil
}

// NOTE: deletes the index if exists and re indexes
func InitRestaurantsElastic(
	elastic *elasticsearch.Client,
	restaurants []models.Restaurant,
	config types.RestaurantsConfig,
) error {
	ctx := context.Background()
	bulk, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:      config.IndexName,
		Client:     elastic,
		NumWorkers: 5,
	})
	if err != nil {
		return err
	}

	for _, item := range restaurants {
		body, err := json.Marshal(item)
		if err != nil {
			return err
		}
		err = bulk.Add(ctx, esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: item.ID,
			Body:       bytes.NewReader(body),
		})
		if err != nil {
			return err
		}
	}

	err = bulk.Close(ctx)
	if err != nil {
		return err
	}

	return nil
}

type RestaurantReader struct {
	cvsReader *csv.Reader
	fid       *os.File
}

// 0:id 1:name 2:address 3:phone 4:Longitude 5:Latitude
func (restaurantReader *RestaurantReader) Read(tar *models.Restaurant) error {
	slices, err := restaurantReader.cvsReader.Read()
	if nil != err {
		return err
	}
	rid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	created := time.Now()
	updated := time.Now()
	tar.ID = rid.String()
	tar.Name = slices[1]
	tar.Address = slices[2]
	tar.Phone = slices[3]
	tar.CreatedAt = created
	tar.UpdatedAt = updated
	tar.Location.CreatedAt = created
	tar.Location.UpdatedAt = updated
	tar.Location.Lat, err = strconv.ParseFloat(slices[4], 64)
	tar.Location.ID = uuid.New().String()
	if err != nil {
		return err
	}
	tar.Location.Lon, err = strconv.ParseFloat(slices[5], 64)
	if err != nil {
		return err
	}
	return nil
}

func newRestaurantsReader(path string) (*RestaurantReader, error) {
	reader := RestaurantReader{}
	fid, err := os.Open(path)
	if err != nil {
		return &reader, err
	}
	reader.fid = fid
	scvReader := csv.NewReader(fid)
	scvReader.Comma = rune('\t')
	reader.cvsReader = scvReader
	scvReader.Read()
	return &reader, nil
}

func ReadAllRestaurants(
	filePath string,
) ([]models.Restaurant, error) {
	reader, err := newRestaurantsReader(filePath)
	defer reader.fid.Close()
	if err != nil {
		return nil, err
	}

	restaurant := models.Restaurant{}
	res := []models.Restaurant{}

	for {
		err := reader.Read(&restaurant)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		res = append(res, restaurant)
	}

	return res, nil
}
