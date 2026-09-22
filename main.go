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

// Входная точка приложения, где создается экземпляр бота и запускается цикл обработки апдейтов
func main() {
	err := godotenv.Load() // Загрузка переменных окружения из файла .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TELEGRAM_BOT_TOKEN") // Получение токена бота из переменных окружения

	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}

	telegramBot, err := telegram.NewBotAPI(token) // Создание нового экземпляра бота с использованием токена
	if err != nil {
		log.Fatal(err)
	}

	telegramBot.Debug = bot.Debug // Флаг для включения режима отладки бота, определенный в пакете bot

	log.Printf("Bot started: @%s", telegramBot.Self.UserName)

	updateConfig := telegram.NewUpdate(0)     // Создание конфигурации для получения обновлений от Telegram
	updateConfig.Timeout = bot.PollingTimeout // Установка таймаута для получения обновлений, определенного в пакете bot

	updates := telegramBot.GetUpdatesChan(updateConfig) // Получение канала обновлений от Telegram с использованием конфигурации

	for update := range updates {
		handlers.HandleUpdate(telegramBot, update) // Обработка каждого обновления с использованием функции HandleUpdate из пакета handlers
	}
}
