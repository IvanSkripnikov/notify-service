package models

type Database struct {
	Address  string
	Port     int
	User     string
	Password string
	DB       string
}

type Redis struct {
	Address  string
	Port     string
	Password string
	DB       int
	Stream   string
}
