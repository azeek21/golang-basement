package repository

import "gorm.io/gorm"

type LoggingRepository struct {
	DB *gorm.DB
}

func NewLoggingRepository(db *gorm.DB) LoggingRepository {
	return LoggingRepository{
		DB: db,
	}
}

func (r LoggingRepository) CreateLog() string {
	return "implement me"
}

func (r LoggingRepository) GetLogById(string) string {
	return "implement me"
}
