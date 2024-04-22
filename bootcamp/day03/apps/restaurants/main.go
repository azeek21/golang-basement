package restaurants

import (
	"github.com/elastic/go-elasticsearch/v8"

	"server/core/env"
)

func InitRestaurants(ec *elasticsearch.Client, env *env.Env) error {
	if err := loadRestaurants(ec, env.InitialDataPath, env.RestaurantsIndex); err != nil {
		return err
	}
	return nil
}
