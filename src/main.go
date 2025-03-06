package main

import (
	"context"

	"notify-service/events"
	"notify-service/helpers"
	"notify-service/httphandler"
	"notify-service/models"

	logger "github.com/IvanSkripnikov/go-logger"
	migrator "github.com/IvanSkripnikov/go-migrator"
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
	migrator.CreateTables(helpers.DB)

	// инициализация REST-api
	httphandler.InitHTTPServer()

	logger.Info("Service started")
}
