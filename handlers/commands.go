package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// функция которая отправяет что-то на определенную команду можешь через switch
func HandleCommands(ctx *HandlerContext) {
	if ctx.Update.Message.IsCommand() && ctx.Update.Message.Command() == "start" {
		HandleStart(ctx)
	}
}

func HandleStart(ctx *HandlerContext) {
	log.Printf("%#v", ctx.UserInfo) //next to db
	msg := tgbotapi.NewMessage(ctx.UserInfo.ChatID, "Прив)")
	_, err := ctx.Bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
