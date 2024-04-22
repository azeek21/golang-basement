package types

import "server/core/env"

type RestaurantsConfig struct {
	DataPath  string `env:"RESTAURANTS_PATH"`
	IndexName string `env:"RESTAURANTS_INDEX"`
}

type ElasticEnv struct {
	Url              string `env:"EC_URL"`
	User             string `env:"EC_USERNAME"`
	Pass             string `env:"EC_PASSWORD"`
	CaCert           string `env:"EC_CERT"`
	InitialDataPath  string `env:"DATA_PATH"`
	RestaurantsIndex string `env:"EC_INDEX_RESTAURANTS"`
}

type PostgresDbConfig struct {
	Host     string `env:"PG_HOST"`
	User     string `env:"PG_USER"`
	Name     string `env:"PG_DB_NAME"`
	Password string `env:"PG_PASSWORD"`
}

type ServerConfig struct {
	Port string `env:"SRV_PORT"`
}

type Tokens struct {
	Key    string `env:"KEY"`
	Secret string `env:"SECRET"`
}

type GlobalConfig struct {
	PostgresDbConfig
	RestaurantsConfig
	ElasticEnv
	ServerConfig
	Tokens
}

var GLOABAL_CONFIG = GlobalConfig{}

func LoadGlobalConfig() error {
	restaurantConfig, err := env.LoadEnv(RestaurantsConfig{})
	if err != nil {
		return err
	}
	elasticConfig, err := env.LoadEnv(ElasticEnv{})
	if err != nil {
		return err
	}
	serverConfig, err := env.LoadEnv(ServerConfig{})
	if err != nil {
		return err
	}
	postgresConfig, err := env.LoadEnv(PostgresDbConfig{})
	if err != nil {
		return err
	}

	tokens, err := env.LoadEnv(Tokens{})
	if err != nil {
		return err
	}
	GLOABAL_CONFIG.PostgresDbConfig = postgresConfig
	GLOABAL_CONFIG.ElasticEnv = elasticConfig
	GLOABAL_CONFIG.ServerConfig = serverConfig
	GLOABAL_CONFIG.RestaurantsConfig = restaurantConfig
	GLOABAL_CONFIG.Tokens = tokens
	return nil
}
