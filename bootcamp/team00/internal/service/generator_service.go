package service

import (
	"math/rand"

	"github.com/google/uuid"
)

type GeneratorService struct {
}

func (s GeneratorService) float64Range(from, to float64) float64 {
	return from + rand.Float64()*(to-from)
}

func NewGeneratorService() GeneratorService {
	return GeneratorService{}
}

func (s GeneratorService) Uuid() string {
	return uuid.New().String()
}

func (s GeneratorService) Mean(from, to float64) float64 {
	return s.float64Range(from, to)
}

func (s GeneratorService) Deviation(from, to float64) float64 {
	return s.float64Range(from, to)
}

func (s GeneratorService) Frequency(mean, deviation float64) float64 {
	// NormFloat64 returns a normally distributed float64 in the range
	// [-math.MaxFloat64, +math.MaxFloat64] with
	// standard normal distribution (mean = 0, stddev = 1)
	// from the default Source.
	// To produce a different normal distribution, callers can
	// adjust the output using:
	//
	//	sample = NormFloat64() * desiredStdDev + desiredMean
	return rand.NormFloat64()*deviation + mean
}
