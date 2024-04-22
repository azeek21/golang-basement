package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"server/types"
)

func NewPostgresDb(config types.PostgresDbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.User,
		config.Name,
		config.Password,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
