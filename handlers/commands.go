package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// функция которая отправяет что-то на определенную команду можешь через switch
func HandleCommands(ctx *HandlerContext) {
	if ctx.Update.Message.IsCommand() {
		switch ctx.Update.Message.Command() {
		case "start":
			HandleStart(ctx)
		}
	} else {
		msg := tgbotapi.NewMessage(ctx.UserInfo.ChatID, "Введите правильную команду!")
		_, err := ctx.Bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func HandleStart(ctx *HandlerContext) {
	log.Printf("%#v", ctx.UserInfo) //next to db
	text := `👋 Добро пожаловать в ScheduleBot!

	Я помогу вам всегда быть в курсе расписания — без лишних поисков и скриншотов.

	Что я умею:
	📅 Показывать расписание на сегодня, завтра и неделю
	🔔 Напоминать о начале пар/встреч
	🔄 Сообщать об изменениях и заменах
	⚙️ Настраиваться под чётные/нечётные недели

	Чтобы начать, выберите свою группу 👇`

	msg := tgbotapi.NewMessage(ctx.UserInfo.ChatID, text)
	_, err := ctx.Bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
