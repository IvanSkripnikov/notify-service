package helpers

import (
	"net/http"
	"strings"

	"notify-service/models"

	"github.com/IvanSkripnikov/go-gormdb"
)

func GetNotificationsList(w http.ResponseWriter, _ *http.Request) {
	category := "/v1/notifications/list"
	var notifications []models.Notification

	db := gormdb.GetClient(models.ServiceDatabase)
	err := db.Find(&notifications).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": notifications,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetNotifications(w http.ResponseWriter, r *http.Request) {
	category := "/v1/notifications/get"
	var notifications []models.Notification

	UserID, _ := getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if UserID == 0 {
		FormatResponse(w, http.StatusUnprocessableEntity, category)
		return
	}

	db := gormdb.GetClient(models.ServiceDatabase)
	err := db.Where("user_id = ?", UserID).Find(&notifications).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": notifications,
	}
	SendResponse(w, data, category, http.StatusOK)
}
