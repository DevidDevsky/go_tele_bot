/*
   ╔══════════════════════════════════════════════════════════════╗
   ║                                                              ║
   ║        ██████╗  ██████╗     ████████╗███████╗██╗     ███████╗ ║
   ║       ██╔════╝ ██╔═══██╗    ╚══██╔══╝██╔════╝██║     ██╔════╝ ║
   ║       ██║  ███╗██║   ██║       ██║   █████╗  ██║     █████╗   ║
   ║       ██║   ██║██║   ██║       ██║   ██╔══╝  ██║     ██╔══╝   ║
   ║       ╚██████╔╝╚██████╔╝       ██║   ███████╗███████╗███████╗ ║
   ║        ╚═════╝  ╚═════╝        ╚═╝   ╚══════╝╚══════╝╚══════╝ ║
   ║                                                              ║
   ║                 TELEGRAM EDUCATION BOT                      ║
   ║                                                              ║
   ║                 devskyyy × pupsmane                         ║
   ║                                                              ║
   ╚══════════════════════════════════════════════════════════════╝

   Educational multi-messenger bot project.
   Stack: Go • PostgreSQL • Telegram
*/

package main

import (
	"go_tele_bot/bot"
	"go_tele_bot/handlers"
	"log"
	"os"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}

	telegramBot, err := telegram.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	telegramBot.Debug = bot.Debug

	log.Printf("Bot started: @%s", telegramBot.Self.UserName)

	updateConfig := telegram.NewUpdate(0)
	updateConfig.Timeout = bot.PollingTimeout

	updates := telegramBot.GetUpdatesChan(updateConfig)

	for update := range updates {
		handlers.HandleUpdate(telegramBot, update)
	}
}
