package repository

import (
	"randomaliens/internal/models"

	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return SessionRepository{
		db: db,
	}
}

func (r SessionRepository) Create(sesison *models.Session) (string, error) {
	if sesison == nil {
		sesison = &models.Session{}
	}
	if err := r.db.Create(sesison).Error; err != nil {
		return "", err
	}
	return sesison.ID, nil
}

func (r SessionRepository) GetById(sessiodId string) (*models.Session, error) {
	panic("not implemented") // TODO: Implement
}

func (r SessionRepository) GetAllAnomaliesOfSession(sessionId string) ([]models.Anomaly, error) {
	panic("not implemented") // TODO: Implement
}

func (r SessionRepository) Update(session *models.Session) error {
	err := r.db.Save(session).Error
	return err

}
