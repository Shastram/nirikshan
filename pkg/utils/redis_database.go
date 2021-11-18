package utils

import "github.com/go-redis/redis/v8"

func RedisDatabaseConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPassword,
		DB:       0,
	})
	return rdb
}
