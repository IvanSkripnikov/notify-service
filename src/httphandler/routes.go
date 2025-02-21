package httphandler

import (
	"net/http"
	"regexp"

	"notify-service/controllers"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

var routes = []route{
	// system
	newRoute(http.MethodGet, "/health", controllers.HealthCheck),
	// notifications
	newRoute(http.MethodGet, "/v1/notifications/list", controllers.GetNotificationsListV1),
	newRoute(http.MethodGet, "/v1/notifications/get/([0-9]+)", controllers.GetNotificationsV1),
}
