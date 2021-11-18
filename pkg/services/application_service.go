package services

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"nirikshan-backend/pkg/records"
	"nirikshan-backend/pkg/siteconfigs"
	"nirikshan-backend/pkg/user"
)

type ApplicationService interface {
	userService
	siteConfigService
	userRecordService
	redisService
}

type applicationService struct {
	userRepository       user.Repository
	siteConfigRepository siteconfigs.Repository
	userRecordRepository records.Repository
	db                   *mongo.Database
	rdb                  *redis.Client
}

// NewService creates a new instance of this service
func NewService(userRepo user.Repository, siteConfigsRepo siteconfigs.
	Repository, userRecordRepository records.Repository,
	db *mongo.Database, rdb *redis.Client) ApplicationService {
	return &applicationService{
		userRepository:       userRepo,
		siteConfigRepository: siteConfigsRepo,
		userRecordRepository: userRecordRepository,
		db:                   db,
		rdb:                  rdb,
	}
}
