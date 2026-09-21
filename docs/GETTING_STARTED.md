# 🚀 Руководство по началу работы

Добро пожаловать в **AI Student Assistant**! Это руководство поможет вам быстро начать разработку.

---

## 📋 Содержание

- [Требования](#-требования)
- [Установка](#-установка)
- [Настройка](#-настройка)
- [Запуск](#-запуск)
- [Первые шаги](#-первые-шаги)
- [Структура проекта](#-структура-проекта)
- [Полезные команды](#-полезные-команды)
- [Устранение проблем](#-устранение-проблем)

---

## 🔧 Требования

Перед началом убедитесь, что у вас установлено:

### Обязательные
- **Go** 1.27 или выше
  - Проверка: `go version`
  - [Скачать Go](https://go.dev/dl/)

- **Git**
  - Проверка: `git --version`
  - [Скачать Git](https://git-scm.com/downloads)

### Необходимо для работы
- **Telegram аккаунт**
- **Telegram Bot Token** от [@BotFather](https://t.me/BotFather)

### Рекомендуемые
- **VS Code** с расширением Go
- **GoLand** / другая Go IDE
- **Postman** для тестирования API (опционально)

---

## 📥 Установка

### 1. Клонирование репозитория

```bash
# HTTPS
git clone https://github.com/yourusername/go_tele_bot.git

# SSH (если настроен)
git clone git@github.com:yourusername/go_tele_bot.git

# Перейти в директорию
cd go_tele_bot
```

### 2. Установка зависимостей

```bash
go mod download
```

Проверка установленных зависимостей:
```bash
go mod verify
```

---

## ⚙️ Настройка

### 1. Создание Telegram бота

Если у вас ещё нет бота:

1. Откройте Telegram и найдите [@BotFather](https://t.me/BotFather)
2. Отправьте команду `/newbot`
3. Следуйте инструкциям:
   - Введите имя бота (например: `My Student Assistant`)
   - Введите username бота (должен заканчиваться на `bot`, например: `my_student_assistant_bot`)
4. Сохраните полученный **токен** (выглядит примерно так: `1234567890:ABCdefGHIjklMNOpqrsTUVwxyz`)

### 2. Настройка .env файла

Создайте `.env` файл на основе примера:

```bash
cp .env.example .env
```

Откройте `.env` в текстовом редакторе и добавьте ваш токен:

```env
TELEGRAM_BOT_TOKEN=ваш_токен_здесь
```

⚠️ **Важно:** Не коммитьте `.env` в Git! Он уже добавлен в `.gitignore`.

### 3. Проверка конфигурации

```bash
# Проверить, что токен загружается
go run main.go
```

Вы должны увидеть:
```
Bot started: @your_bot_username
```

---

## ▶️ Запуск

### Запуск в режиме разработки

```bash
go run main.go
```

### Сборка и запуск

```bash
# Сборка
go build -o bot

# Запуск (Linux/macOS)
./bot

# Запуск (Windows)
bot.exe
```

### Запуск с параметрами

```bash
# С указанием .env файла
go run main.go -env=.env.development

# С debug режимом
DEBUG=true go run main.go
```

---

## 🎯 Первые шаги

### 1. Проверка работы бота

После запуска откройте Telegram и:

1. Найдите вашего бота по username
2. Нажмите **Start** или отправьте `/start`
3. Бот должен ответить

### 2. Просмотр логов

Логи появляются в консоли:

```
2026/09/19 10:30:00 Bot started: @my_student_assistant_bot
```

### 3. Остановка бота

Нажмите `Ctrl+C` в терминале.

---

## 📁 Структура проекта

```
go_tele_bot/
├── 📂 bot/              # Конфигурация бота
│   └── settings.go      # Параметры (Debug, Timeout, и т.д.)
│
├── 📂 handlers/         # Обработчики событий
│   ├── handler.go       # Главный роутер
│   ├── commands.go      # /start, /help и другие команды
│   ├── messages.go      # Текстовые сообщения от пользователей
│   └── callbacks.go     # Обработка нажатий на inline кнопки
│
├── 📂 keyboards/        # Inline и Reply клавиатуры
│   └── keyboards.go     # Генерация клавиатур
│
├── 📂 docs/             # Документация
│   └── GETTING_STARTED.md  # Это руководство
│
├── 📄 main.go           # Точка входа
├── 📄 .env.example      # Пример конфигурации
├── 📄 .env              # Ваша конфигурация (не в Git!)
├── 📄 go.mod            # Зависимости Go
├── 📄 go.sum            # Checksums зависимостей
├── 📄 README.md         # Главная документация
└── 📄 CONTRIBUTING.md   # Гайд для контрибьюторов
```

### Описание компонентов

| Файл/Папка | Назначение |
|------------|------------|
| `main.go` | Инициализация бота, загрузка `.env`, запуск polling |
| `bot/settings.go` | Конфигурация: debug режим, timeout для polling |
| `handlers/handler.go` | Маршрутизация updates от Telegram |
| `handlers/commands.go` | Обработка команд типа `/start`, `/help` |
| `handlers/messages.go` | Обработка обычных текстовых сообщений |
| `handlers/callbacks.go` | Обработка callback_query от inline кнопок |
| `keyboards/keyboards.go` | Создание клавиатур для взаимодействия |

---

## 💻 Полезные команды

### Go команды

```bash
# Форматирование кода
go fmt ./...

# Проверка на ошибки
go vet ./...

# Сборка проекта
go build -o bot

# Запуск
go run main.go

# Установка зависимостей
go mod download

# Обновление зависимостей
go get -u ./...
go mod tidy

# Очистка кэша
go clean -cache
```

### Git команды

```bash
# Проверка статуса
git status

# Создание новой ветки
git checkout -b feature/my-feature

# Коммит изменений
git add .
git commit -m "feat: add new feature"

# Отправка на GitHub
git push origin feature/my-feature

# Обновление из main
git fetch origin
git rebase origin/main
```

### Разработка

```bash
# Просмотр логов в реальном времени
go run main.go 2>&1 | tee bot.log

# Запуск с автоперезагрузкой (с Air)
air

# Проверка версии Go
go version

# Список зависимостей
go list -m all
```

---

## 🔧 Устранение проблем

### Проблема: "TELEGRAM_BOT_TOKEN is not set"

**Причина:** Токен не найден в `.env` файле.

**Решение:**
1. Убедитесь, что файл `.env` существует в корне проекта
2. Проверьте, что токен указан правильно:
   ```env
   TELEGRAM_BOT_TOKEN=1234567890:ABCdefGHI...
   ```
3. Перезапустите бота

### Проблема: "Error loading .env file"

**Причина:** Файл `.env` не найден.

**Решение:**
```bash
cp .env.example .env
# Затем отредактируйте .env и добавьте токен
```

### Проблема: Bot не отвечает в Telegram

**Возможные причины:**
1. Неверный токен - проверьте токен в `.env`
2. Бот не запущен - убедитесь, что `go run main.go` работает
3. Проблемы с сетью - проверьте интернет-соединение

**Решение:**
```bash
# Перезапустите бота с debug режимом
# В bot/settings.go установите:
# Debug = true
```

### Проблема: "package not found"

**Причина:** Зависимости не установлены.

**Решение:**
```bash
go mod download
go mod tidy
```

### Проблема: Изменения в коде не применяются

**Решение:**
1. Остановите бота (`Ctrl+C`)
2. Очистите кэш: `go clean -cache`
3. Пересоберите: `go build`
4. Запустите снова: `go run main.go`

---

## 📚 Следующие шаги

После успешного запуска изучите:

1. 📖 [CONTRIBUTING.md](../CONTRIBUTING.md) - как добавлять новый функционал
2. 📖 [README.md](../README.md) - общий обзор проекта
3. 🔍 Исследуйте код в `handlers/` - примеры обработчиков
4. 💡 Добавьте свою первую команду в `handlers/commands.go`
5. ⌨️ Создайте свою клавиатуру в `keyboards/keyboards.go`

---

## 🆘 Нужна помощь?

- 📖 Читайте [документацию проекта](../README.md)
- 🐛 Создайте [Issue](https://github.com/yourusername/go_tele_bot/issues)
- 💬 Напишите в Telegram
- 📧 Отправьте email команде

---

## 🔗 Полезные ссылки

- [Go Documentation](https://go.dev/doc/)
- [Telegram Bot API](https://core.telegram.org/bots/api)
- [go-telegram-bot-api](https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api/v5)
- [Effective Go](https://go.dev/doc/effective_go)

---

<div align="center">

**Удачи в разработке!** 🚀

[⬆ Вернуться наверх](#-руководство-по-началу-работы)

</div>
