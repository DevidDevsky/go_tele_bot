package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
		Bot:      bot,
		Update:   update,
		UserInfo: userInfo,
	}

	if update.Message.IsCommand() && update.Message.Command() == "start" {
		HandleStart(&ctx)
	}
}
