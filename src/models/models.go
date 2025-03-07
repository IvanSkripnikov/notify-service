package models

const ServiceDatabase = "NotificationService"

type Redis struct {
	Address  string
	Port     string
	Password string
	DB       int
	Stream   string
}
