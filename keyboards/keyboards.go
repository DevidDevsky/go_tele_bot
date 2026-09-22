package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartKeyboard() tgbotapi.InlineKeyboardMarkup {
	scheduleButton := tgbotapi.NewInlineKeyboardButtonData("Расписание", "schedule")
	landmarkButton := tgbotapi.NewInlineKeyboardButtonData("Ориентир", "landmark")

	row := []tgbotapi.InlineKeyboardButton{scheduleButton, landmarkButton}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)
	return keyboard
}
