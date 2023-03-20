package main

import (
	"awesomeProject1/pkg/telegram"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6062992339:AAEniV0HSE7vEaHXszIcSoDMSSabdXCknfI")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	telegramBot := telegram.Newbot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
