package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model

	ID                   string `gorm:"primaryKey"`
	ReceivedRecordsCount int
	FrequencySum         float64
	StandartDeviation    float64
	Mean                 float64
}

type Anomaly struct {
	gorm.Model

	ID         string `gorm:"primaryKey"`
	SessionID  string
	Session    Session
	ReceivedAt string
	Difference float64
}

// NOTE sessions should get their id from server. So there's no uuid generating precreate functin for session

func (s *Anomaly) BeforeCreate(db *gorm.DB) error {
	newUuid := uuid.New()
	db.Statement.SetColumn("ID", newUuid.String())
	return nil
}
