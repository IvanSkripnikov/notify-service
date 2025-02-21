package controllers

import (
	"net/http"

	"notify-service/helpers"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.HealthCheck(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/health")
	}
}
