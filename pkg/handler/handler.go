package handler

import (
	"net/http"

	"github.com/dmitry-bakeev/flight-schedule/pkg/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	{
		api.HandleFunc("/flight", h.CreateMultiple).Methods(http.MethodPost)
	}
	return router
}
