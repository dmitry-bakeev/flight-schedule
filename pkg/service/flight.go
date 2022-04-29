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

func (s *FlightService) FilterFromCity(city string) ([]*models.Flight, error) {
	return s.repo.FilterFromCity(city)
}

func (s *FlightService) FilterToCity(city string) ([]*models.Flight, error) {
	return s.repo.FilterToCity(city)
}
