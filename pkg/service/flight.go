package service

import (
	"fmt"

	"github.com/dmitry-bakeev/flight-schedule/pkg/models"
	"github.com/dmitry-bakeev/flight-schedule/pkg/repository"
	"github.com/dmitry-bakeev/flight-schedule/sort"
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

func (s *FlightService) OrderByNumberFlight(desc bool) ([]*models.Flight, error) {
	flights, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	sorted := make(chan []*models.Flight)
	go sort.MergeSort(flights, sorted)

	result := <-sorted
	close(sorted)

	if desc {
		sort.Reverse(result)
	}

	return result, nil
}

func (s *FlightService) OrderByFromCity(desc bool) ([]*models.Flight, error) {
	return s.repo.Order("from_city", desc)
}

func (s *FlightService) OrderByTimeFromCity(desc bool) ([]*models.Flight, error) {
	return s.repo.Order("time_from_city", desc)
}

func (s *FlightService) OrderByToCity(desc bool) ([]*models.Flight, error) {
	return s.repo.Order("to_city", desc)
}

func (s *FlightService) OrderByTimeToCity(desc bool) ([]*models.Flight, error) {
	return s.repo.Order("time_to_city", desc)
}
