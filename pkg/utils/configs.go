package utils

import (
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	JwtSecret              = os.Getenv("JWT_SECRET")
	DBUrl                  = os.Getenv("DB_SERVER")
	DBUser                 = os.Getenv("DB_USER")
	DBPassword             = os.Getenv("DB_PASSWORD")
	JWTExpiryDuration      = getEnvAsInt("JWT_EXPIRY_MINS", 1440)
	DBName                 = "niriskhan"
	Port                   = ":3000"
	UserCollection         = "users"
	UsernameField          = "username"
	SiteConfigCollection   = "site_configs"
	SiteNameField          = "site_name"
	UserRecordsCollection  = "user_dump_records"
	TelegramBotToken       = os.Getenv("TELEGRAM_BOT_TOKEN")
	TelegramUser           = getEnvAsInt("TELEGRAM_USER", 0)
	NirikshanBackendGithub = "https://github.com/Shastram/nirikshan-backend"
	RedisPassword          = ""
	RedisAddr              = os.Getenv("REDIS_SERVER")
	DdosCountLimit         = 4
	DdosExpirationTime     = time.Second * 20

	PasswordEncryptionCost = bcrypt.DefaultCost
)

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := os.Getenv(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
