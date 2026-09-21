package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type UserInfo struct {
UserID    int64
ChatID    int64
Username  string
FirstName string
LastName  string
}

type HandlerContext struct {
Bot      *tgbotapi.BotAPI
Update   tgbotapi.Update
UserInfo UserInfo
}
