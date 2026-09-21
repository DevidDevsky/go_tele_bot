package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Написать функцию через switch которая определяет тип и относительно типа через case вызывает функцию в 1) commands(если сообщение == /и слово) 2) обработка кнопки если нажимают на кнопку 3) обработка текста который пишет чел мяу мяу

type UserInfo struct {
	UserID    int64
	ChatID    int64
	Username  string
	FirstName string
	LastName  string
}

type HandlerContext struct {
	Update   tgbotapi.Update
	UserInfo UserInfo
	Bot      *tgbotapi.BotAPI
}

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	userID := update.Message.From.ID
	chatID := update.Message.Chat.ID
	username := update.Message.From.UserName
	firstName := update.Message.From.FirstName
	lastName := update.Message.From.LastName

	userInfo := UserInfo{
		UserID:    userID,
		ChatID:    chatID,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
	}

	ctx := HandlerContext{
		Update:   update,
		UserInfo: userInfo,
		Bot:      bot,
	}

	HandleCommands(&ctx)
}
