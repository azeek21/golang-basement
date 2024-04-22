package restaurants

import (
	"errors"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"

	"server/core/env"
)

var RestaurantMappings = `
{
  "settings": {
    "number_of_shards": 1,
    "max_result_window": 20000
  },
  "mappings": {
    "properties": {
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

func CreateRestaurantsIndex(client *elasticsearch.Client, env *env.Env) error {
	res, err := client.Indices.Delete([]string{env.RestaurantsIndex})

	if err != nil || res.IsError() {
		if res.StatusCode != 404 {
			return errors.New(
				fmt.Sprintf(
					"error [deleting old index: %s]: %v\n",
					env.RestaurantsIndex,
					res.String(),
				),
			)
		} else {
			return err
		}
	}

	res, err = client.Indices.Create(
		env.RestaurantsIndex,
		client.Indices.Create.WithBody(strings.NewReader(RestaurantMappings)),
	)

	if res.IsError() {
		return errors.New(fmt.Sprintf("error [creating restaurants index]: %v\n", res.Body))
	}

	if err != nil {
		return err
	}

	return nil
}
