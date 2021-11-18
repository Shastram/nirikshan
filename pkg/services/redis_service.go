package services

import (
	"context"
	"github.com/go-redis/redis/v8"
	"nirikshan-backend/pkg/utils"
)

type redisService interface {
	GetKey(key string) (string, error)
	PutData(key string, value string) error
}

func (a applicationService) GetKey(key string) (string, error) {
	val2, err := a.rdb.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", redis.Nil
	} else if err != nil {
		return "", err
	}
	return val2, nil
}

func (a applicationService) PutData(key string, value string) error {
	err := a.rdb.Set(context.Background(), key, value, utils.DdosExpirationTime).Err()
	if err != nil {
		return err
	}
	return nil
}
