package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartKeyboard() tgbotapi.InlineKeyboardMarkup {
	scheduleButton := tgbotapi.NewInlineKeyboardButtonData("Расписание", "schedule")
	newsButton := tgbotapi.NewInlineKeyboardButtonData("Новости", "news")

	row := []tgbotapi.InlineKeyboardButton{scheduleButton, newsButton}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)
	return keyboard
}
