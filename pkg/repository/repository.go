package repository

import (
	"github.com/dmitry-bakeev/flight-schedule/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Flight interface {
	CreateMultiple(flights []*models.Flight) error
	GetAll() ([]*models.Flight, error)
	FilterFromCity(city string) ([]*models.Flight, error)
	FilterToCity(city string) ([]*models.Flight, error)
}

type Repository struct {
	Flight
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Flight: NewFlightPostgres(db),
	}
}
