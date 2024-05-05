package repository

// All actions with db
import (
	"randomaliens/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Logging interface {
	CreateLog() string
	GetLogById(string) string
}

type Session interface {
	Create(*models.Session) (string, error)
	GetById(sessiodId string) (*models.Session, error)
	GetAllAnomaliesOfSession(sessionId string) ([]models.Anomaly, error)
	Update(*models.Session) error
}
type Anomaly interface {
	GetById(id string) (*models.Anomaly, error)
	Create(anomalRecord *models.Anomaly) (string, error)
}

type Repository struct {
	Logging
	Anomaly
	Session
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Logging: NewLoggingRepository(db),
		Anomaly: NewAnomalyRepository(db),
		Session: NewSessionRepository(db),
	}
}

func NewPostgresDb() (*gorm.DB, error) {
	dsn := "host=localhost user=azeek dbname=goday03 password=admin sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
