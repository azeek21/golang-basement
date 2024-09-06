package database

import "errors"

var (
	DB_ERR_NOT_FOUND      = errors.New("not found")
	DB_ERR_ALREADY_EXISTS = errors.New("an entry with this key aleready exists")
)

type DB interface {
	Create(key, value string) error
	Read(key string) (string, error)
	Update(key, value string) error
	Delete(key string) error
}

type db struct {
	storage map[string]string
}

func CreateDB() DB {
	return &db{
		storage: make(map[string]string),
	}
}

func (d db) Read(key string) (string, error) {
	res, ok := d.storage[key]
	if !ok {
		return res, DB_ERR_NOT_FOUND
	}
	return res, nil
}

func (d db) Create(key, value string) error {
	_, exits := d.storage[key]
	if exits {
		return DB_ERR_ALREADY_EXISTS
	}

	d.storage[key] = value

	return nil
}

func (d db) Update(key, value string) error {
	_, exists := d.storage[key]
	if !exists {
		return DB_ERR_NOT_FOUND
	}

	d.storage[key] = value

	return nil
}

func (d db) Delete(key string) error {
	delete(d.storage, key)
	return nil
}
