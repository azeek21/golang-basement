package service

import (
	"log"
	"math"
	"randomaliens/internal/grpc/transmitter"
	"randomaliens/internal/models"
	"randomaliens/internal/repository"
)

type SessionService struct {
	repo *repository.Repository
}

type AnomalySession struct {
	sumSq     float64
	threshold float64
	models.Session
	sessionService *SessionService
}

func newAnomlaySession(service *SessionService) AnomalySession {
	return AnomalySession{
		sessionService: service,
	}
}

func NewSessionSevice(repo *repository.Repository) SessionService {
	return SessionService{
		repo: repo,
	}
}

func (s SessionService) NewSession(id string, threshold float64) (*AnomalySession, error) {
	newSession := AnomalySession{
		threshold: threshold,
		Session: models.Session{
			ID: id,
		},
	}
	_, err := s.repo.Session.Create(&newSession.Session)
	if err != nil {
		return nil, err
	}

	return &newSession, nil
}

func (s SessionService) SaveSession(session *AnomalySession) error {
	if session == nil {
		return nil
	}
	return s.repo.Session.Update(&session.Session)
}

func (s SessionService) SaveAnomaly(record *transmitter.Record, sesison *AnomalySession) error {
	s.repo.Anomaly.Create(&models.Anomaly{
		Session:    sesison.Session,
		ReceivedAt: *record.Timestamp,
		Difference: sesison.getZScore(record),
	})
	return nil
}

func (s *AnomalySession) Record(record *transmitter.Record) error {
	s.Session.ReceivedRecordsCount += 1
	s.Session.FrequencySum += *record.Frequency
	s.Session.Mean = s.FrequencySum / float64(s.ReceivedRecordsCount)
	s.sumSq += math.Pow(*record.Frequency, 2)
	s.Session.StandartDeviation = math.Sqrt((s.sumSq - s.FrequencySum*s.FrequencySum/float64(s.ReceivedRecordsCount)) / float64(s.ReceivedRecordsCount-1))

	log.Printf("s: mean: %v, sd: %v, c %v\n",
		s.Mean,
		s.StandartDeviation,
		s.ReceivedRecordsCount,
	)
	return nil
}

func (s *AnomalySession) IsAnomaly(record *transmitter.Record) bool {
	zSore := math.Abs(s.getZScore(record))
	isAnom := zSore > s.threshold
	log.Printf("Checking anomaly z: %v, thresh: %v, res: %v", zSore, s.threshold, isAnom)
	return isAnom
}

func (s *AnomalySession) getZScore(record *transmitter.Record) float64 {
	return (*record.Frequency - s.Mean) / s.StandartDeviation
}
