package repository

import (
	"github.com/elastic/go-elasticsearch/v8"

	"server/types"
)

func GetElasticConfig(env *types.ElasticEnv) elasticsearch.Config {
	return elasticsearch.Config{
		Addresses:              []string{env.Url},
		Password:               env.Pass,
		Username:               env.User,
		CertificateFingerprint: env.CaCert,
	}
}

func NewElasticClient(env *types.ElasticEnv) (*elasticsearch.Client, error) {
	config := GetElasticConfig(env)

	client, err := elasticsearch.NewClient(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
