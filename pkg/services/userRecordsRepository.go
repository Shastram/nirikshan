package services

import (
	"nirikshan-backend/pkg/entities"
)

type userRecordService interface {
	GetDump(siteName string) (*[]entities.UserRecords, error)
	CreateDump(configs *entities.UserRecords) error
}

func (a applicationService) GetDump(siteName string) (*[]entities.UserRecords, error) {
	return a.userRecordRepository.GetDump(siteName)
}

func (a applicationService) CreateDump(configs *entities.UserRecords) error {
	return a.userRecordRepository.CreateDump(configs)
}
