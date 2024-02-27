package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hanacore/config"
	send "hanacore/helper/message"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.Text == "/start" { // If we got a message
			msg := send.CreateMessage(update.Message.From.ID, update.Message.MessageID, update.Message.Text)

			FirstName := update.Message.From.FirstName
			MessageText := fmt.Sprint("Hello, " + FirstName + "! Howdy?")
			msg.Text = MessageText
			send.SendMessage(bot, msg)
			continue
		}
	}
}
