package main

import (
	"context"

	"notify-service/events"
	"notify-service/helpers"
	"notify-service/httphandler"
	"notify-service/models"

	"github.com/IvanSkripnikov/go-gormdb"
	logger "github.com/IvanSkripnikov/go-logger"
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
	helpers.InitDatabase(config.Database)

	// настройка коннекта к redis
	bus := events.MakeBus()
	go helpers.Listen(bus)
	helpers.InitRedis(context.Background(), config.Redis)

	go helpers.ListenStream(helpers.HandleMessage, bus.Error)

	// выполнение миграций
	migrationModels := models.GetModels()
	gormdb.ApplyMigrationsForClient(models.ServiceDatabase, migrationModels...)
	//migrator.CreateTables(helpers.DB)

	// инициализация REST-api
	httphandler.InitHTTPServer()

	logger.Info("Service started")
}
