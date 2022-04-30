package repository

import (
	"fmt"

	"github.com/dmitry-bakeev/flight-schedule/pkg/models"
	"github.com/jmoiron/sqlx"
)

type FlightPostgres struct {
	db *sqlx.DB
}

func NewFlightPostgres(db *sqlx.DB) *FlightPostgres {
	return &FlightPostgres{
		db: db,
	}
}

func (r *FlightPostgres) CreateMultiple(flights []*models.Flight) error {
	query := fmt.Sprintf("INSERT INTO %s(number_flight, from_city, time_from_city, to_city, time_to_city) VALUES (:number_flight, :from_city, :time_from_city, :to_city, :time_to_city)", flightTable)
	_, err := r.db.NamedExec(query, flights)

	return err
}

func (r *FlightPostgres) GetAll() ([]*models.Flight, error) {
	var result []*models.Flight

	query := fmt.Sprintf("SELECT number_flight, from_city, time_from_city, to_city, time_to_city FROM %s", flightTable)
	err := r.db.Select(&result, query)

	return result, err
}

func (r *FlightPostgres) FilterFromCity(city string) ([]*models.Flight, error) {
	var result []*models.Flight

	query := fmt.Sprintf("SELECT number_flight, from_city, time_from_city, to_city, time_to_city FROM %s WHERE from_city = $1", flightTable)
	err := r.db.Select(&result, query, city)

	return result, err
}

func (r *FlightPostgres) FilterToCity(city string) ([]*models.Flight, error) {
	var result []*models.Flight

	query := fmt.Sprintf("SELECT number_flight, from_city, time_from_city, to_city, time_to_city FROM %s WHERE to_city = $1", flightTable)
	err := r.db.Select(&result, query, city)

	return result, err
}

func (r *FlightPostgres) Order(field string, desc bool) ([]*models.Flight, error) {
	var result []*models.Flight
	var order string

	if desc {
		order = "DESC"
	} else {
		order = "ASC"
	}

	query := fmt.Sprintf("SELECT number_flight, from_city, time_from_city, to_city, time_to_city FROM %s ORDER BY %s %s", flightTable, field, order)
	err := r.db.Select(&result, query)

	return result, err
}
