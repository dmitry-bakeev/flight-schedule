package service

import (
	"github.com/dmitry-bakeev/flight-schedule/pkg/models"
	"github.com/dmitry-bakeev/flight-schedule/pkg/repository"
)

type Flight interface {
	CreateMultiple(flights []*models.Flight) error
}

type Service struct {
	Flight
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Flight: NewFlightService(repo.Flight),
	}
}
