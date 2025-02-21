package controllers

import (
	"net/http"

	"notify-service/helpers"
)

func GetNotificationsListV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetNotificationsList(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/notifications/list")
	}
}

func GetNotificationsV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetNotifications(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/notifications/get")
	}
}
