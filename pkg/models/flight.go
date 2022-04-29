package models

import (
	"errors"
	"strings"
	"time"
)

type Flight struct {
	Id           int
	NumberFlight int
	FromCity     string
	TimeFromCity time.Time
	ToCity       string
	TimeToCity   time.Time
}

func (f *Flight) Validate() error {
	fields := make([]string, 0)
	if f.NumberFlight == 0 {
		fields = append(fields, "number_flight is required")
	}
	if f.FromCity == "" {
		fields = append(fields, "from_city is required")
	}
	if f.TimeFromCity.IsZero() {
		fields = append(fields, "time_from_city is required")
	}
	if f.ToCity == "" {
		fields = append(fields, "to_city is required")
	}
	if f.TimeToCity.IsZero() {
		fields = append(fields, "time_to_city is required")
	}

	if len(fields) != 0 {
		return errors.New(strings.Join(fields, "; "))
	}

	return nil
}
