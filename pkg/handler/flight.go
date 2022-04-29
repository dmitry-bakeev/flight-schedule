package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dmitry-bakeev/flight-schedule/pkg/models"
)

const (
	filterFromCity = "filter_from_city"
	filterToCity   = "filter_to_city"
)

type Input struct {
	Objects []*models.Flight `json:"objects"`
}

func (h *Handler) CreateMultiple(w http.ResponseWriter, r *http.Request) {
	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Flight.CreateMultiple(input.Objects)
	if err != nil {
		newErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, r, http.StatusCreated, map[string]interface{}{
		"added": "done",
	})
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	fromCity := query.Get(filterFromCity)
	if fromCity != "" {
		result, err := h.services.FilterFromCity(fromCity)
		if err != nil {
			newErrorResponse(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		newResponse(w, r, http.StatusCreated, map[string]interface{}{
			"objects": result,
		})
	}

	toCity := query.Get(filterToCity)
	if toCity != "" {
		result, err := h.services.FilterToCity(toCity)
		if err != nil {
			newErrorResponse(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		newResponse(w, r, http.StatusCreated, map[string]interface{}{
			"objects": result,
		})
	}
}
