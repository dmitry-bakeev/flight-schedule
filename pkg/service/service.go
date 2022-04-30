package service

import (
	"github.com/dmitry-bakeev/flight-schedule/pkg/models"
	"github.com/dmitry-bakeev/flight-schedule/pkg/repository"
)

type Flight interface {
	CreateMultiple(flights []*models.Flight) error
	FilterFromCity(city string) ([]*models.Flight, error)
	FilterToCity(city string) ([]*models.Flight, error)
	OrderByNumberFlight(desc bool) ([]*models.Flight, error)
	OrderByFromCity(desc bool) ([]*models.Flight, error)
	OrderByTimeFromCity(desc bool) ([]*models.Flight, error)
	OrderByToCity(desc bool) ([]*models.Flight, error)
	OrderByTimeToCity(desc bool) ([]*models.Flight, error)
}

type Service struct {
	Flight
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Flight: NewFlightService(repo.Flight),
	}
}
