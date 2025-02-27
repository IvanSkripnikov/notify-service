package models

import (
	"os"
	"strconv"
)

type Config struct {
	Database Database
	Redis    Redis
}

func LoadConfig() (*Config, error) {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))
	if err != nil {
		return nil, err
	}

	return &Config{
		Database: Database{
			Address:  os.Getenv("DB_ADDRESS"),
			Port:     dbPort,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DB:       os.Getenv("DB_NAME"),
		},
		Redis: Redis{
			Address:  os.Getenv("REDIS_ADDRESS"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       redisDB,
			Stream:   os.Getenv("REDIS_STREAM"),
		},
	}, nil
}

func GetRequiredVariables() []string {
	return []string{
		// Обязательные переменные окружения для подключения к БД сервиса
		"DB_ADDRESS", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME",

		// Обязательные переменные окружения для подключения к redis
		"REDIS_ADDRESS", "REDIS_PORT", "REDIS_PASSWORD", "REDIS_DB_NUMBER", "REDIS_STREAM",
	}
}
