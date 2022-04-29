package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

func newResponse(w http.ResponseWriter, r *http.Request, statusCode int, message interface{}) {
	w.WriteHeader(statusCode)

	result, err := json.Marshal(Response{
		Status:  "ok",
		Message: message,
	})
	if err != nil {
		log.Print(err.Error())
	}

	w.Write(result)
}

func newErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	w.WriteHeader(statusCode)

	result, err := json.Marshal(Response{
		Status:  "error",
		Message: message,
	})
	if err != nil {
		log.Print(err.Error())
	}

	w.Write(result)
}
