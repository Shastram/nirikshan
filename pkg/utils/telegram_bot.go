package utils

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func InitialiseTelegramBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(TelegramBotToken)
	if err != nil {
		log.Error(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	go func(bot *tgbotapi.BotAPI) {
		updates := bot.GetUpdatesChan(u)
		for update := range updates {
			if update.Message == nil {
				continue
			}

			if !update.Message.IsCommand() {
				continue
			}
			var msg tgbotapi.MessageConfig
			switch update.Message.Command() {
			case "start":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID,
					startTemplate(update.Message.From.UserName))
				break
			case "getid":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, strconv.FormatInt(update.Message.Chat.ID, 10))
				break
			}
			bot.Send(msg)
		}
	}(bot)
	return bot
}

func startTemplate(username string) string {
	var startTemplate = fmt.Sprintf(
		"Hello %s ! Welcome to Nirakshan !\n\n"+
			"Steps to use Nirakshan: \n"+
			"1. use /getid to generate your unique id \n\n"+
			"2. Go to %s for more info \n\n"+
			"3. Paste your unique id in the env file of your backend"+
			" \n\n"+
			"4. That's it, you will be notified of any DDOS attacks , "+
			"enjoy :D \n\n\n", username, NirikshanBackendGithub)
	return startTemplate
}
