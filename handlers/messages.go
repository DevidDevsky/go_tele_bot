package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Для ключей вида "26"/"27" лучше использовать map.
type first_group map[string]string

func (g first_group) reply(text string) string {
	switch text {
	case "26":
		return g["26"]
	case "27":
		return g["27"]
	default:
		return ""
	}
}

func messageHandler(ctx *HandlerContext) {
	if ctx.Update.Message != nil {
		switch ctx.Update.Message.Text {
		case "26":
			reply := first_group{}

			reply["26"] = "Расписание для группы 26: \n1. Математика\n2. Физика\n3. Химия"
			msg := tgbotapi.NewMessage(ctx.UserInfo.ChatID, reply.reply("26"))
			ctx.Bot.Send(msg)
		}
	}
}
