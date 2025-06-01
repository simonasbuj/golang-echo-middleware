package service

import (
	"time"
	"math"
)


type Service struct {}

func New() *Service {
	return &Service{}
}

// Calculate the number of days until January 1, 2027
func (s *Service) DaysUntil2027() int {
	targetDate := time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC)
	currentDate := time.Now().UTC()
	daysUntil := targetDate.Sub(currentDate).Hours() / 24

	return int(math.Ceil(daysUntil))
}

