package helpers

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"notify-service/events"
	"notify-service/models"

	"github.com/IvanSkripnikov/go-logger"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	cont        context.Context
	stream      string
)

// Init Инициализация подключения к Redis.
func InitRedis(ctx context.Context, config models.Redis) {
	if _, err := strconv.Atoi(config.Port); err != nil {
		logger.Fatalf("Failed to parse on Redis port. Error: %v", err)
	}

	address := net.JoinHostPort(config.Address, config.Port)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: config.Password,
		DB:       config.DB,
	})
	cont = ctx
	stream = config.Stream
	logger.Info("Redis initialized")
}

// Listen Прослушивать сообщения в каналах.
func Listen(bus events.EventBus) {
	for {
		select {
		case err := <-bus.Error:
			logger.Error(err.Error())
		}
	}
}

// ListenStream Прослушивание стрима Redis.
func ListenStream(handler func(redis.XMessage), errCh chan<- error) {
	logger.Info("Listening stream...")
	lastId := "0"
	for {
		result, err := redisClient.XRead(cont, &redis.XReadArgs{
			Count:   100,
			Block:   0,
			Streams: []string{stream, lastId},
		}).Result()

		if err != nil {
			logger.Errorf("Cant execute XRead command. Error: %v", err)
			errCh <- err
			return
		}

		messages := result[0].Messages
		countMessages := len(messages)

		if countMessages > 0 {
			logger.Debugf("XRead iteration from ID: %s. New messages: %d", lastId, countMessages)
			lastId = messages[countMessages-1].ID
		}

		for _, message := range messages {
			handler(message)
		}
	}
}

func HandleMessage(message redis.XMessage) {
	var err error
	category := fmt.Sprint(message.Values["category"])
	if category != "deal" && category != "loyalty" {
		logger.Warning("Unknown message category")
		return
	}

	logger.Info("New message: " + message.ID)
	title := fmt.Sprint(message.Values["title"])
	description := fmt.Sprint(message.Values["description"])
	userID := fmt.Sprint(message.Values["user"])

	var notification models.Notification
	notification.Title = title
	notification.Description = description
	notification.UserID, err = strconv.Atoi(userID)
	if err != nil {
		logger.Errorf("Cant convert UserID to int %v", err)
	}
	notification.Created = int(GetCurrentTimestamp())

	logger.Debug(fmt.Sprintf("Message %s value: %v", message.ID, notification))

	// записываем сообщение в БД
	err = GormDB.Create(&notification).Error
	if err != nil {
		logger.Errorf("Cant create notification message %v", err)
	}

	// удаляем сообщение из стрима
	if errDel := DeleteMessage(message.ID); errDel != nil {
		logger.Error("Cant delete message " + message.ID + " from redis stream")
	}
	MessagesTotal.Inc()
}

// DeleteMessage Удалить сообщение из стрима.
func DeleteMessage(id string) error {
	_, err := redisClient.XDel(cont, stream, id).Result()
	if err != nil {
		logger.Errorf("Cant execute XDel command on message %s. Error: %v", id, err)
		return err
	}

	logger.Infof("Message %s deleted from Redis", id)
	return nil
}
