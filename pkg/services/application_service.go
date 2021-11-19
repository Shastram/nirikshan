package services

import (
	"nirikshan-backend/pkg/records"
	"nirikshan-backend/pkg/siteconfigs"
	"nirikshan-backend/pkg/user"

	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApplicationService interface {
	userService
	siteConfigService
	userRecordService
	redisService
	telegramBotService
}

type applicationService struct {
	userRepository       user.Repository
	siteConfigRepository siteconfigs.Repository
	userRecordRepository records.Repository
	db                   *mongo.Database
	rdb                  *redis.Client
	telegramBotApi       *tgbotapi.BotAPI
}

// NewService creates a new instance of this service
func NewService(userRepo user.Repository, siteConfigsRepo siteconfigs.
	Repository, userRecordRepository records.Repository,
	db *mongo.Database, rdb *redis.Client,
	teleBot *tgbotapi.BotAPI) ApplicationService {
	return &applicationService{
		userRepository:       userRepo,
		siteConfigRepository: siteConfigsRepo,
		userRecordRepository: userRecordRepository,
		db:                   db,
		rdb:                  rdb,
		telegramBotApi:       teleBot,
	}
}
