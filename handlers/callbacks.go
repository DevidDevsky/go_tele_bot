package handlers

import (
	"log"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleCallback(bot *telegram.BotAPI, update telegram.Update) {
	switch update.CallbackQuery.Data {
	case "schedule":
		handleScheduleButton(bot, update)

	case "news":
		handleNewsButton(bot, update)

	case "rooms":
		handleRoomsButton(bot, update)
	}

	_, err := bot.Request(
		telegram.NewCallback(update.CallbackQuery.ID, ""),
	)

	if err != nil {
		log.Println(err)
	}
}

func handleScheduleButton(bot *telegram.BotAPI, update telegram.Update) {
	msg := telegram.NewMessage(
		update.CallbackQuery.Message.Chat.ID,
		"Вы выбрали кнопку 'Расписание'.",
	)

	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func handleNewsButton(bot *telegram.BotAPI, update telegram.Update) {
	msg := telegram.NewMessage(
		update.CallbackQuery.Message.Chat.ID,
		"Вы выбрали кнопку 'Новости'.",
	)

	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func handleRoomsButton(bot *telegram.BotAPI, update telegram.Update) {
	msg := telegram.NewMessage(
		update.CallbackQuery.Message.Chat.ID,
		"Вы выбрали кнопку 'Аудитории'.",
	)

	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
