package helpers

import (
	"net/http"
	"strconv"
	"strings"

	"notify-service/models"

	logger "github.com/IvanSkripnikov/go-logger"
)

func GetNotificationsList(w http.ResponseWriter, _ *http.Request) {
	category := "/v1/notifications/list"
	var notifications []models.Notification

	query := "SELECT id, user_id, title, description, created FROM notifications WHERE id > 0"
	rows, err := DB.Query(query)
	if err != nil {
		logger.Error(err.Error())
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	for rows.Next() {
		notification := models.Notification{}
		if err = rows.Scan(&notification.ID, &notification.UserID, &notification.Title, &notification.Description, &notification.Created); err != nil {
			logger.Error(err.Error())
			continue
		}
		notifications = append(notifications, notification)
	}

	data := ResponseData{
		"data": notifications,
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

	if !isExists("SELECT * FROM notifications WHERE user_id = ?", UserID) {
		FormatResponse(w, http.StatusNotFound, category)
		return
	}

	query := "SELECT id, user_id, title, description, created FROM notifications WHERE user_id = " + strconv.Itoa(UserID)
	rows, err := DB.Query(query)
	if err != nil {
		logger.Error(err.Error())
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	for rows.Next() {
		notification := models.Notification{}
		if err = rows.Scan(&notification.ID, &notification.UserID, &notification.Title, &notification.Description, &notification.Created); err != nil {
			logger.Error(err.Error())
			continue
		}
		notifications = append(notifications, notification)
	}

	data := ResponseData{
		"data": notifications,
	}
	SendResponse(w, data, category, http.StatusOK)
}
