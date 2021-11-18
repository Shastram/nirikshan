package services

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"nirikshan-backend/pkg/utils"
)

type telegramBotService interface {
	SendMessage(message string) error
}

func (a applicationService) SendMessage(message string) error {
	msg := tgbotapi.NewMessage(int64(utils.TelegramUser), message)
	_, err := a.telegramBotApi.Send(msg)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
