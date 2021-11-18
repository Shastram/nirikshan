package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"nirikshan-backend/pkg/records"
	"nirikshan-backend/pkg/siteconfigs"
	"nirikshan-backend/pkg/user"
)

type ApplicationService interface {
	userService
	siteConfigService
	userRecordService
}

type applicationService struct {
	userRepository       user.Repository
	siteConfigRepository siteconfigs.Repository
	userRecordRepository records.Repository
	db                   *mongo.Database
}

// NewService creates a new instance of this service
func NewService(userRepo user.Repository, siteConfigsRepo siteconfigs.
	Repository, userRecordRepository records.Repository,
	db *mongo.Database) ApplicationService {
	return &applicationService{
		userRepository:       userRepo,
		siteConfigRepository: siteConfigsRepo,
		userRecordRepository: userRecordRepository,
		db:                   db,
	}
}
