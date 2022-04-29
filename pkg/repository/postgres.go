package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	flightTable = "flight"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func NewPostgresDB(cfg *PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"pgx",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s database=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
		),
	)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
