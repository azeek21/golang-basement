package repository

import (
	"randomaliens/internal/models"

	"gorm.io/gorm"
)

type AnomalyRepository struct {
	db *gorm.DB
}

func NewAnomalyRepository(db *gorm.DB) AnomalyRepository {
	return AnomalyRepository{
		db: db,
	}
}

func (r AnomalyRepository) Create(anomalRecord *models.Anomaly) (string, error) {
	if err := r.db.Create(anomalRecord).Error; err != nil {
		return "", err
	}
	return anomalRecord.ID, nil
}
func (r AnomalyRepository) GetById(id string) (*models.Anomaly, error) {
	res := &models.Anomaly{ID: id}
	if err := r.db.Preload("Session").Find(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
