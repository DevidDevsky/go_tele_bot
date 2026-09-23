package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Структура для хранения информации о пользователе
type UserInfo struct {
	UserID    int64
	ChatID    int64
	Username  string
	FirstName string
	LastName  string
}

// Структура для хранения контекста обработки апдейта
type HandlerContext struct {
	Update   tgbotapi.Update
	UserInfo UserInfo
	Bot      *tgbotapi.BotAPI
}

// Функция которая определяет какой тип апдейта пришел и вызывает соответствующую функцию для обработки вызывается в main.go
func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var userInfo UserInfo

	if update.Message != nil {
		userInfo = UserInfo{
			UserID:    update.Message.From.ID,
			ChatID:    update.Message.Chat.ID,
			Username:  update.Message.From.UserName,
			FirstName: update.Message.From.FirstName,
			LastName:  update.Message.From.LastName,
		}
	}

	if update.CallbackQuery != nil {
		userInfo = UserInfo{
			UserID:    update.CallbackQuery.From.ID,
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			Username:  update.CallbackQuery.From.UserName,
			FirstName: update.CallbackQuery.From.FirstName,
			LastName:  update.CallbackQuery.From.LastName,
		}
	}

	ctx := HandlerContext{
		Update:   update,
		UserInfo: userInfo,
		Bot:      bot,
	}

	if update.Message != nil {
		if update.Message.IsCommand() {
			HandleCommands(&ctx)
			return
		} else {
			messageHandler(&ctx)
			return
		}
	}

	if update.CallbackQuery != nil {
		HandleCallbackQuery(&ctx)
		return
	}
}
