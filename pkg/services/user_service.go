package services

import (
	"context"
	"nirikshan-backend/pkg/entities"
)

type userService interface {
	GetUser(uid string) (*entities.User, error)
	GetUsers() (*[]entities.User, error)
	FindUsersByUID(uid []string) (*[]entities.User, error)
	FindUserByUsername(username string) (*entities.User, error)
	CheckPasswordHash(hash, password string) error
	UpdatePassword(userPassword *entities.UserPassword, isAdminBeingReset bool, ctx context.Context) error
	CreateUser(user *entities.User, ctx context.Context) error
	UpdateUser(user *entities.User, ctx context.Context) error
}

func (a applicationService) GetUser(uid string) (*entities.User, error) {
	return a.userRepository.GetUser(uid)
}

func (a applicationService) GetUsers() (*[]entities.User, error) {
	return a.userRepository.GetUsers()
}

func (a applicationService) FindUsersByUID(uid []string) (*[]entities.User, error) {
	return a.userRepository.FindUsersByUID(uid)
}

func (a applicationService) FindUserByUsername(username string) (*entities.User, error) {
	return a.userRepository.FindUserByUsername(username)
}

func (a applicationService) CheckPasswordHash(hash, password string) error {
	return a.userRepository.CheckPasswordHash(hash, password)
}

func (a applicationService) UpdatePassword(userPassword *entities.UserPassword, isAdminBeingReset bool, ctx context.Context) error {
	return a.userRepository.UpdatePassword(userPassword, isAdminBeingReset, ctx)
}

func (a applicationService) CreateUser(user *entities.User, ctx context.Context) error {
	return a.userRepository.CreateUser(user, ctx)
}

func (a applicationService) UpdateUser(user *entities.User, ctx context.Context) error {
	return a.UpdateUser(user, ctx)
}
