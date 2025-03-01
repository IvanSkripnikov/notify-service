package main

import (
	"context"

	"notify-service/events"
	"notify-service/helpers"
	"notify-service/httphandler"
	"notify-service/logger"
	"notify-service/models"
)

func main() {
	logger.Debug("Service starting")

	// регистрация общих метрик
	helpers.RegisterCommonMetrics()

	// настройка всех конфигов
	config, err := models.LoadConfig()
	if err != nil {
		logger.Fatalf("Config error: %v", err)
	}

	// настройка коннекта к БД
	_, err = helpers.InitDataBase(config.Database)
	if err != nil {
		logger.Fatalf("Cant initialize DB: %v", err)
	}

	// настройка коннекта к redis
	bus := events.MakeBus()
	go helpers.Listen(bus)
	helpers.InitRedis(context.Background(), config.Redis)

	go helpers.ListenStream(helpers.HandleMessage, bus.Error)

	// выполнение миграций
	helpers.CreateTables()

	// инициализация REST-api
	httphandler.InitHTTPServer()

	logger.Info("Service started")
}
