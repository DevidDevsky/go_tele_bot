package main

import (
	"go_tele_bot/bot"
	"go_tele_bot/handlers"
	"log"
	"os"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set in the environment variables")
	}

	telegramBot, err := telegram.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Failed to create Telegram bot: %v", err)
	}

	telegramBot.Debug = bot.Debug
	log.Printf("Authorized on account %s", telegramBot.Self.UserName)

	updateConfig := telegram.NewUpdate(0)
	updateConfig.Timeout = bot.PollingTimeout
	updates := telegramBot.GetUpdatesChan(updateConfig)
	for update := range updates {
		handlers.HandleUpdate(telegramBot, update)
	}

}
