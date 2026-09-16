package main

// log — стандартная библиотека для логирования
// tgbotapi(alias) "url библиотеки" — библиотека для работы с Telegram Bot API
import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Передаём токен и создаём объект бота
	bot, err := tgbotapi.NewBotAPI("Token")

	// Обработка ошибки. Panic — вывод ошибки и остановка программы
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true                                          // Включаем дополнительную отладочную информацию
	log.Printf("Authorized on account %s", bot.Self.UserName) // Self — информация о самом боте, UserName — username бота

	// Настройка бота
	u := tgbotapi.NewUpdate(0)       // Настройка получения обновлений
	u.Timeout = 60                   // Время ожидания новых обновлений при long polling
	updates := bot.GetUpdatesChan(u) // Получаем канал с обновлениями от Telegram

	// Берём очередное событие из канала updates
	for update := range updates {
		// Проверяем, содержит ли Update обычное сообщение
		if update.Message != nil {
			// Выводим username пользователя и текст сообщения
			log.Printf(
				"[%s] %s",
				update.Message.From.UserName,
				update.Message.Text,
			)

			msg := tgbotapi.NewMessage(
				update.Message.Chat.ID,
				update.Message.Text,
			) // Создаём сообщение: первый аргумент — Chat ID, второй — текст сообщения

			msg.ReplyToMessageID = update.Message.MessageID // Делаем наше сообщение ответом на сообщение пользователя

			_, err := bot.Send(msg) // Отправляем сообщение

			// Проверяем ошибку отправки
			if err != nil {
				log.Println(err)
			}
		}
	}
}
