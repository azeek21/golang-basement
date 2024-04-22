package initializers

import (
	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

type Initializer struct {
	db      *gorm.DB
	elastic *elasticsearch.TypedClient
}

func NewInitializer(db *gorm.DB, elastic *elasticsearch.TypedClient) *Initializer {
	return &Initializer{
		db:      db,
		elastic: elastic,
	}
}
