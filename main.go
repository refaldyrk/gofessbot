package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	CHANNEL_ID = "YOUR CHANNEL ID"
	ADMIN_ID   = "YOUR_ID"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("APIFESS"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.Chat.Type == "private" {
				checkPrefix := strings.Split(update.Message.Text, " ")
				log.Printf("[NEW MESSAGE]")

				if checkPrefix[0] == "#fess" {
					fesser := strings.Join(checkPrefix[1:len(checkPrefix)], " ")

					msg := tgbotapi.NewMessageToChannel(CHANNEL_ID, fesser)

					_, err := bot.Send(msg)
					if err != nil {
						log.Printf("[ERROR] %x", err)
					}
				}
			}

			if update.Message.Chat.Type == "channel" {
				log.Println("[NEW MESSAGE FROM CHANNEL]")
			}
		}
	}
}
