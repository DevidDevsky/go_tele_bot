package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Функция которая определяет какой тип апдейта пришел и вызывает соответствующую функцию для обработки
func HandleCallbackQuery(ctx *HandlerContext) {
	if ctx.Update.CallbackQuery != nil {
		switch ctx.Update.CallbackQuery.Data {
		case "schedule":
			HandleSchedule(ctx)
		case "news":
			HandleNews(ctx)
		default:
			HandleUnknownCallback(ctx)
		}
	}
}

// Функция которая обрабатывает нажатие на кнопку "Расписание"
func HandleSchedule(ctx *HandlerContext) {
	msg := tgbotapi.NewMessage(ctx.UserInfo.ChatID, "Введите номер группы!")

	ctx.Bot.Send(msg)
}

// Функция которая обрабатывает нажатие на кнопку "Новости"
func HandleNews(ctx *HandlerContext) {
	msg := tgbotapi.NewMessage(ctx.UserInfo.ChatID, "Вы нажали кнопку новости!")
	ctx.Bot.Send(msg)
}

// Функция которая обрабатывает неизвестные кнопки
func HandleUnknownCallback(ctx *HandlerContext) {
	msg := tgbotapi.NewMessage(ctx.UserInfo.ChatID, "Неизвестная кнопка!")
	ctx.Bot.Send(msg)
}
