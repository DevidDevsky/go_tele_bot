# ⚡ Быстрая шпаргалка

Краткий справочник для быстрого старта и ежедневной разработки.

---

## 🚀 Быстрый старт

```bash
# Клонировать
git clone https://github.com/yourusername/go_tele_bot.git
cd go_tele_bot

# Установить зависимости
go mod download

# Создать .env
cp .env.example .env
# Добавить токен в .env

# Запустить
go run main.go
```

---

## 📝 Основные команды

### Go команды

```bash
# Запуск
go run main.go

# Сборка
go build -o bot

# Форматирование
go fmt ./...

# Проверка
go vet ./...

# Тесты (когда будут)
go test ./...

# Обновление зависимостей
go get -u ./...
go mod tidy
```

### Git команды

```bash
# Создать ветку
git checkout -b feature/название

# Коммит
git add .
git commit -m "feat: описание"

# Пуш
git push origin feature/название

# Обновиться из main
git fetch origin
git rebase origin/main
```

---

## 📁 Структура проекта

```
go_tele_bot/
├── bot/              # Настройки бота
├── handlers/         # Обработчики событий
├── keyboards/        # Клавиатуры
├── docs/             # Документация
├── main.go           # Точка входа
└── .env              # Конфигурация (НЕ в Git!)
```

---

## 🔧 Типичные задачи

### Добавить новую команду

**Файл:** `handlers/commands.go`

```go
func HandleMyCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    msg := tgbotapi.NewMessage(message.Chat.ID, "Ответ")
    bot.Send(msg)
}
```

**Регистрация:** `handlers/handler.go`

```go
case "/mycommand":
    handlers.HandleMyCommand(bot, update.Message)
```

### Добавить клавиатуру

**Файл:** `keyboards/keyboards.go`

```go
func GetMyKeyboard() tgbotapi.InlineKeyboardMarkup {
    return tgbotapi.NewInlineKeyboardMarkup(
        tgbotapi.NewInlineKeyboardRow(
            tgbotapi.NewInlineKeyboardButtonData("Кнопка", "callback_data"),
        ),
    )
}
```

### Обработать callback

**Файл:** `handlers/callbacks.go`

```go
func HandleCallback(bot *tgbotapi.BotAPI, query *tgbotapi.CallbackQuery) {
    switch query.Data {
    case "callback_data":
        // Ваша логика
    }
}
```

---

## 💡 Соглашения

### Commit сообщения

```bash
feat: новая функция
fix: исправление бага
docs: документация
refactor: рефакторинг
test: тесты
chore: технические изменения
```

### Названия веток

```bash
feature/add-schedule    # новая функция
fix/callback-error      # исправление
refactor/handlers       # рефакторинг
docs/update-readme      # документация
```

### Именование в Go

```go
// Экспортируемые (публичные) - с большой буквы
func HandleStart() {}
type UserProfile struct {}

// Неэкспортируемые (приватные) - с маленькой
func processMessage() {}
type internalState struct {}
```

---

## 🐛 Отладка

### Включить debug режим

**Файл:** `bot/settings.go`

```go
const (
    Debug = true  // Установить true
)
```

### Логирование

```go
import "log"

log.Printf("Debug: %v", data)
log.Printf("Error: %v", err)
```

### Проверка токена

```bash
# Проверить, что токен загружается
cat .env
# Должен показать: TELEGRAM_BOT_TOKEN=...
```

---

## ⚠️ Частые ошибки

### "TELEGRAM_BOT_TOKEN is not set"

```bash
# Решение:
cp .env.example .env
# Отредактировать .env и добавить токен
```

### "package not found"

```bash
# Решение:
go mod download
go mod tidy
```

### Бот не отвечает

```bash
# Проверить:
1. Бот запущен? (go run main.go)
2. Токен правильный? (проверить в .env)
3. Интернет работает?
4. Debug режим включен? (в bot/settings.go)
```

### Изменения не применяются

```bash
# Решение:
Ctrl+C  # Остановить бота
go clean -cache
go run main.go  # Запустить снова
```

---

## 📚 Полезные ссылки

- [README](../README.md) - обзор проекта
- [CONTRIBUTING](../CONTRIBUTING.md) - как контрибутить
- [GETTING_STARTED](./GETTING_STARTED.md) - детальный гайд
- [PROJECT_INFO](./PROJECT_INFO.md) - информация о проекте

### Внешние ресурсы

- [Go by Example](https://gobyexample.com/)
- [Telegram Bot API](https://core.telegram.org/bots/api)
- [go-telegram-bot-api Docs](https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api/v5)

---

## 🎯 Быстрые примеры

### Отправить сообщение

```go
msg := tgbotapi.NewMessage(chatID, "Текст")
bot.Send(msg)
```

### Отправить с клавиатурой

```go
msg := tgbotapi.NewMessage(chatID, "Текст")
msg.ReplyMarkup = keyboards.GetMainKeyboard()
bot.Send(msg)
```

### Ответить на callback

```go
callback := tgbotapi.NewCallback(query.ID, "Текст")
bot.Request(callback)
```

### Редактировать сообщение

```go
edit := tgbotapi.NewEditMessageText(
    chatID,
    messageID,
    "Новый текст",
)
bot.Send(edit)
```

---

## 🔐 Безопасность

### ✅ Делать

- Использовать `.env` для секретов
- Добавить `.env` в `.gitignore`
- Валидировать пользовательский ввод
- Обрабатывать ошибки

### ❌ Не делать

- Коммитить `.env` в Git
- Хардкодить токены в коде
- Игнорировать ошибки
- Логировать секреты

---

## 📊 Статус разработки

```
✅ Инициализация проекта
✅ Telegram Bot API
✅ Базовые handlers
🔄 Клавиатуры (в процессе)
⏳ База данных (планируется)
⏳ Расписание (планируется)
⏳ AI функции (планируется)
```

---

## 📞 Помощь

**Проблемы?**
- 🔍 Поищите в [Issues](https://github.com/yourusername/go_tele_bot/issues)
- 🆕 Создайте [новый Issue](https://github.com/yourusername/go_tele_bot/issues/new)
- 💬 Напишите в Telegram

---

<div align="center">

**Полезные горячие клавиши VS Code**

| Команда | Действие |
|---------|----------|
| `Cmd/Ctrl + P` | Быстрый поиск файла |
| `Cmd/Ctrl + Shift + F` | Поиск в проекте |
| `Cmd/Ctrl + /` | Комментарий |
| `F5` | Запуск дебагера |
| `Shift + Alt + F` | Форматирование |

---

*Всё, что нужно для быстрого старта!* ⚡

</div>
