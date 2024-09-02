package repository

import (
	"replication/database"
	"replication/utils"
)

type Repository interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
}

type repository struct {
	db database.DB
}

func NewReopsitory(db database.DB) Repository {
	return &repository{
		db: db,
	}
}

// FUNCTIONS
func (r repository) Get(key string) (string, error) {
	res, err := r.db.Read(key)
	if err != nil {
		return res, utils.WithPrefix(key, err)
	}
	return res, nil
}

func (r repository) Set(key string, value string) error {
	_, err := r.db.Read(key)

	if err != nil {
		return utils.WithPrefix(key, r.db.Create(key, value))
	}

	return utils.WithPrefix(key, r.db.Update(key, value))
}

func (r repository) Delete(key string) error {
	return utils.WithPrefix(key, r.db.Delete(key))
}
