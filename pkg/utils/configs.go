package utils

import (
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	JwtSecret             = os.Getenv("JWT_SECRET")
	DBUrl                 = os.Getenv("DB_SERVER")
	DBUser                = os.Getenv("DB_USER")
	DBPassword            = os.Getenv("DB_PASSWORD")
	JWTExpiryDuration     = getEnvAsInt("JWT_EXPIRY_MINS", 1440)
	DBName                = "niriskhan"
	Port                  = ":3000"
	UserCollection        = "users"
	UsernameField         = "username"
	SiteConfigCollection  = "site_configs"
	UserRecordsCollection = "user_dump_records"
	RedisPassword         = ""
	RedisAddr             = "localhost:6379"
	DdosCountLimit        = 4
	DdosExpirationTime    = time.Hour * 1

	PasswordEncryptionCost = bcrypt.DefaultCost
)

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := os.Getenv(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
