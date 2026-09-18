package handlers

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(bot *telegram.BotAPI, update telegram.Update) {
	if update.Message != nil {
		handleMessage(bot, update)
		return
	}

	if update.CallbackQuery != nil {
		handleCallback(bot, update)
		return
	}
}
