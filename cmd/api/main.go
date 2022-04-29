package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/dmitry-bakeev/flight-schedule/pkg/handler"
	"github.com/dmitry-bakeev/flight-schedule/pkg/repository"
	"github.com/dmitry-bakeev/flight-schedule/pkg/service"
	"github.com/dmitry-bakeev/flight-schedule/server"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(&repository.PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DATABASE"),
		SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
	})

	if err != nil {
		log.Fatalf("error initializing db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	handler := handler.NewHandler(services)

	srv := server.Server{}

	go func() {
		if err := srv.Run(os.Getenv("RUN_PORT"), handler.InitRoutes()); err != nil {
			log.Fatalf("error running http server: %s", err.Error())
		}
	}()

	log.Printf("server running on port: %s", os.Getenv("RUN_PORT"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Print("server shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error shutting down http server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("error closing db: %s", err.Error())
	}
}
