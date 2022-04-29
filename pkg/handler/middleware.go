package handler

import (
	"fmt"
	"net/http"
	"strings"
)

type ContextKey string

const (
	titleContentHeader = "Content-Type"
	jsonHeader         = "application/json"
)

func (h *Handler) JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(titleContentHeader, jsonHeader)

		header := r.Header.Get(titleContentHeader)
		msgErr := fmt.Sprintf("%s header is not %s", titleContentHeader, jsonHeader)

		if header == "" {
			newErrorResponse(w, r, http.StatusUnsupportedMediaType, msgErr)
			return
		}

		if !strings.Contains(header, jsonHeader) {
			newErrorResponse(w, r, http.StatusUnsupportedMediaType, msgErr)
			return
		}

		next.ServeHTTP(w, r)
	})
}
