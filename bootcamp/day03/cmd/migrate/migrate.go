package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"

	"server/core/env"
	"server/pkg/initializers"
	"server/pkg/repository"
	"server/types"
)

// MIGRATE BOTH POSTGRES AND ELASTIC
func main() {
	elasticEnvConfig := types.ElasticEnv{}
	postgresEnvConfig := types.PostgresDbConfig{}
	restaurantsConfig := types.RestaurantsConfig{}

	// elastic
	elasticEnvConfig, err := env.LoadEnv(elasticEnvConfig)
	if err != nil {
		log.Fatal(err)
		return
	}

	ecfg := repository.GetElasticConfig(&elasticEnvConfig)
	elastic, err := elasticsearch.NewClient(ecfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	// postgres
	postgresEnvConfig, err = env.LoadEnv(postgresEnvConfig)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	postgres, err := repository.NewPostgresDb(postgresEnvConfig)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	restaurantsConfig, err = env.LoadEnv(restaurantsConfig)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	restaurants, err := initializers.ReadAllRestaurants(restaurantsConfig.DataPath)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = initializers.MigrateRestaurantsPostgres(postgres, restaurants)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = initializers.MigrateRestaurantsElastic(elastic, &restaurantsConfig, restaurants)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Println("MIGRATION SUCCESS")
}
