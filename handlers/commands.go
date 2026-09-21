package handlers

import (
"log"

tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleStart(ctx *HandlerContext) {
log.Printf("%#v", ctx.UserInfo) // next to db

msg := tgbotapi.NewMessage(
ctx.UserInfo.ChatID,
"Привет! 👋 Добро пожаловать в бота!",
)

if _, err := ctx.Bot.Send(msg); err != nil {
log.Println(err)
}
}
