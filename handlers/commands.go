package handlers

import (
	"log"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleMessage(bot *telegram.BotAPI, update telegram.Update) {
	if update.Message.IsCommand() {
		handleCommand(bot, update)
		return
	}
}

func handleCommand(bot *telegram.BotAPI, update telegram.Update) {
	switch update.Message.Command() {
	case "start":
		handleStart(bot, update)

	case "help":
		handleHelp(bot, update)

	case "schedule":
		handleSchedule(bot, update)
	}
}

func handleStart(bot *telegram.BotAPI, update telegram.Update) {
	msg := telegram.NewMessage(
		update.Message.Chat.ID,
		"Привет! Я бот-помощник колледжа.",
	)

	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func handleHelp(bot *telegram.BotAPI, update telegram.Update) {
	msg := telegram.NewMessage(
		update.Message.Chat.ID,
		"Я могу помочь вам с расписанием, новостями и информацией о колледже. Используйте команды /schedule, /news и /rooms.",
	)

	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func handleSchedule(bot *telegram.BotAPI, update telegram.Update) {
	msg := telegram.NewMessage(
		update.Message.Chat.ID,
		"Вот ваше расписание на сегодня: ...",
	)

	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
