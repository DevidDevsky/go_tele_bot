package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommands(ctx *HandlerContext) {
	if ctx.Update.Message == nil {
		return
	}

	switch ctx.Update.Message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(
			ctx.UserInfo.ChatID,
			"Привет! 👋 Добро пожаловать в бота!",
		)

		if _, err := ctx.Bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}
