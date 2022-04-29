package service

import (
	"fmt"

	"github.com/dmitry-bakeev/flight-schedule/pkg/models"
	"github.com/dmitry-bakeev/flight-schedule/pkg/repository"
)

type FlightService struct {
	repo repository.Flight
}

func NewFlightService(repo repository.Flight) *FlightService {
	return &FlightService{
		repo: repo,
	}
}

func (s *FlightService) CreateMultiple(flights []*models.Flight) error {
	for i, f := range flights {
		if err := f.Validate(); err != nil {
			return fmt.Errorf("for index: %d error: %w", i, err)
		}
	}

	return s.repo.CreateMultiple(flights)
}
