# 🤝 Руководство по внесению вклада

Спасибо за интерес к проекту **AI Student Assistant**! Мы рады любому вкладу — от исправления опечаток до добавления новых функций.

---

## 📋 Содержание

- [Кодекс поведения](#-кодекс-поведения)
- [Как я могу помочь?](#-как-я-могу-помочь)
- [Процесс разработки](#-процесс-разработки)
- [Структура проекта](#-структура-проекта)
- [Стандарты кода](#-стандарты-кода)
- [Git Workflow](#-git-workflow)
- [Commit Convention](#-commit-convention)
- [Pull Request Process](#-pull-request-process)
- [Тестирование](#-тестирование)
- [Документация](#-документация)

---

## 📜 Кодекс поведения

Участвуя в этом проекте, вы соглашаетесь соблюдать следующие принципы:

- ✅ Будьте уважительны и профессиональны
- ✅ Конструктивная критика приветствуется
- ✅ Помогайте новичкам
- ❌ Не допускается оскорбительное поведение
- ❌ Не допускается дискриминация любого вида

---

## 💡 Как я могу помочь?

### 🐛 Сообщить об ошибке

Нашли баг? Создайте [Issue](https://github.com/yourusername/go_tele_bot/issues/new) с информацией:

```markdown
**Описание ошибки**
Краткое описание проблемы.

**Как воспроизвести**
1. Шаг 1
2. Шаг 2
3. Ошибка появляется...

**Ожидаемое поведение**
Что должно было произойти.

**Скриншоты**
Если применимо.

**Окружение:**
- OS: [например, macOS 14]
- Go версия: [например, 1.27]
- Версия проекта: [например, v0.1.0]
```

### ✨ Предложить новую функцию

Есть идея? Создайте [Feature Request](https://github.com/yourusername/go_tele_bot/issues/new):

```markdown
**Описание функции**
Что вы хотите добавить?

**Мотивация**
Зачем это нужно? Какую проблему это решит?

**Предлагаемое решение**
Как это должно работать?

**Альтернативы**
Рассматривали ли вы другие варианты?
```

### 🔧 Внести код

Хотите что-то исправить или добавить? Отлично! Следуйте [процессу разработки](#-процесс-разработки).

### 📖 Улучшить документацию

Документация важна! Можно:
- Исправить опечатки
- Улучшить примеры
- Добавить пояснения
- Перевести на другие языки

---

## 🔄 Процесс разработки

### 1️⃣ Форк репозитория

Кликните на кнопку **Fork** в правом верхнем углу страницы репозитория.

### 2️⃣ Клонируйте ваш форк

```bash
git clone https://github.com/YOUR_USERNAME/go_tele_bot.git
cd go_tele_bot
```

### 3️⃣ Добавьте upstream remote

```bash
git remote add upstream https://github.com/ORIGINAL_OWNER/go_tele_bot.git
```

### 4️⃣ Создайте ветку

```bash
git checkout -b feature/your-feature-name
```

Правила именования веток:
- `feature/` - новая функциональность
- `fix/` - исправление ошибок
- `refactor/` - рефакторинг
- `docs/` - документация
- `test/` - тесты

Примеры:
```bash
git checkout -b feature/add-schedule-handler
git checkout -b fix/callback-error
git checkout -b docs/update-readme
```

### 5️⃣ Настройте окружение

```bash
# Установите зависимости
go mod download

# Создайте .env файл
cp .env.example .env
# Добавьте ваш TELEGRAM_BOT_TOKEN

# Запустите проект
go run main.go
```

### 6️⃣ Внесите изменения

Пишите код согласно [стандартам](#-стандарты-кода).

### 7️⃣ Тестируйте

```bash
# Запустите тесты (когда они будут добавлены)
go test ./...

# Проверьте форматирование
go fmt ./...

# Проверьте на ошибки
go vet ./...
```

### 8️⃣ Закоммитьте изменения

Следуйте [Commit Convention](#-commit-convention):

```bash
git add .
git commit -m "feat: add schedule handler"
```

### 9️⃣ Синхронизируйте с upstream

```bash
git fetch upstream
git rebase upstream/main
```

### 🔟 Отправьте в ваш форк

```bash
git push origin feature/your-feature-name
```

### 1️⃣1️⃣ Создайте Pull Request

Перейдите на страницу вашего форка на GitHub и нажмите **New Pull Request**.

---

## 📁 Структура проекта

```
go_tele_bot/
├── bot/
│   └── settings.go          # Конфигурация бота
├── handlers/
│   ├── handler.go           # Главный маршрутизатор
│   ├── commands.go          # Обработка команд
│   ├── messages.go          # Обработка сообщений
│   └── callbacks.go         # Обработка callbacks
├── keyboards/
│   └── keyboards.go         # Генерация клавиатур
├── database/               # База данных (планируется)
├── models/                 # Модели данных (планируется)
├── services/               # Бизнес-логика (планируется)
├── utils/                  # Утилиты (планируется)
├── docs/                   # Документация
├── .env.example            # Пример переменных окружения
├── .gitignore              # Git исключения
├── go.mod                  # Go зависимости
├── go.sum                  # Checksums зависимостей
├── main.go                 # Точка входа
├── README.md               # Главная документация
└── CONTRIBUTING.md         # Это руководство
```

### 📝 Ответственность компонентов

| Компонент | Назначение |
|-----------|------------|
| `main.go` | Инициализация приложения и запуск бота |
| `bot/settings.go` | Конфигурация параметров бота |
| `handlers/handler.go` | Маршрутизация Telegram updates |
| `handlers/commands.go` | Логика обработки команд (`/start`, `/help`, etc.) |
| `handlers/messages.go` | Обработка текстовых сообщений |
| `handlers/callbacks.go` | Обработка callback queries от inline кнопок |
| `keyboards/keyboards.go` | Создание inline и reply клавиатур |

---

## 🎨 Стандарты кода

### Go Style Guide

Следуйте [Effective Go](https://go.dev/doc/effective_go) и [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).

### Основные правила

#### ✅ Именование

```go
// Переменные: camelCase
var userName string
var userID int64

// Константы: CamelCase или SCREAMING_SNAKE_CASE
const MaxRetries = 3
const API_TIMEOUT = 30

// Функции: CamelCase
func GetUserSchedule() {}
func processCallback() {}  // приватная функция

// Структуры: CamelCase
type UserProfile struct {
    ID       int64
    Username string
}
```

#### ✅ Форматирование

```bash
# Автоматическое форматирование
go fmt ./...

# Или через gofmt
gofmt -w .
```

#### ✅ Комментарии

```go
// Package handlers содержит обработчики для Telegram bot events.
package handlers

// HandleStart обрабатывает команду /start.
// Возвращает приветственное сообщение с inline keyboard.
func HandleStart(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
    // Создаём клавиатуру
    keyboard := keyboards.GetMainKeyboard()
    
    // Отправляем сообщение
    msg := tgbotapi.NewMessage(message.Chat.ID, "Добро пожаловать!")
    msg.ReplyMarkup = keyboard
    
    _, err := bot.Send(msg)
    return err
}
```

#### ✅ Обработка ошибок

```go
// ✅ Правильно
result, err := someFunction()
if err != nil {
    log.Printf("Error in someFunction: %v", err)
    return err
}

// ❌ Неправильно
result, _ := someFunction()  // не игнорируйте ошибки!
```

#### ✅ Структура файлов

```go
// 1. Package declaration
package handlers

// 2. Imports (стандартные, затем сторонние, затем локальные)
import (
    "log"
    "time"
    
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    
    "go_tele_bot/keyboards"
    "go_tele_bot/models"
)

// 3. Константы
const (
    WelcomeMessage = "Добро пожаловать!"
)

// 4. Переменные
var (
    startTime time.Time
)

// 5. Типы
type Handler struct {
    bot *tgbotapi.BotAPI
}

// 6. Функции
func NewHandler(bot *tgbotapi.BotAPI) *Handler {
    return &Handler{bot: bot}
}
```

---

## 🌿 Git Workflow

### Структура веток

```
main (защищена)
├── develop (основная ветка разработки)
│   ├── feature/add-schedule
│   ├── feature/user-management
│   ├── fix/callback-handler
│   └── refactor/handlers-structure
└── hotfix/critical-bug
```

### Правила работы с ветками

1. **main** — стабильная версия (только через PR)
2. **develop** — активная разработка
3. **feature/** — новая функциональность
4. **fix/** — исправления багов
5. **hotfix/** — срочные исправления для production
6. **refactor/** — рефакторинг без изменения функциональности
7. **docs/** — документация
8. **test/** — добавление тестов

### Примеры названий веток

```bash
feature/add-schedule-handler
feature/inline-keyboard
feature/database-integration
fix/callback-handler-error
fix/timezone-bug
refactor/separate-handlers
refactor/improve-error-handling
docs/update-readme
docs/add-api-documentation
test/add-handler-tests
hotfix/critical-memory-leak
```

---

## 📝 Commit Convention

Используем [Conventional Commits](https://www.conventionalcommits.org/).

### Формат

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Типы коммитов

| Тип | Описание | Пример |
|-----|----------|---------|
| `feat` | Новая функциональность | `feat: add schedule handler` |
| `fix` | Исправление бага | `fix: handle empty callback data` |
| `docs` | Документация | `docs: update README` |
| `style` | Форматирование кода | `style: format handlers package` |
| `refactor` | Рефакторинг | `refactor: separate telegram handlers` |
| `test` | Тесты | `test: add schedule handler tests` |
| `chore` | Технические изменения | `chore: update dependencies` |
| `perf` | Оптимизация производительности | `perf: optimize database queries` |
| `ci` | CI/CD изменения | `ci: add GitHub Actions workflow` |
| `build` | Изменения сборки | `build: update go.mod` |

### Примеры коммитов

#### ✅ Хорошие коммиты

```bash
feat: add schedule handler for student groups
feat(handlers): implement callback routing
fix: resolve nil pointer in message handler
fix(keyboards): correct button layout
docs: add installation instructions
docs(api): document handler functions
refactor: extract keyboard logic to separate package
test: add unit tests for command handlers
chore: update telegram bot api to v5.5.1
```

#### ❌ Плохие коммиты

```bash
update                    # Что обновили?
fix                      # Что исправили?
changes                  # Какие изменения?
new feature              # Какая функция?
asdf                     # Непонятно
final                    # Финальный что?
test                     # Что тестируете?
```

### Детальные коммиты

Для больших изменений используйте body:

```bash
git commit -m "feat: add schedule system

Implement complete schedule management:
- Database schema for schedules
- Handler for schedule requests
- Keyboard navigation
- Date selection logic

Closes #42"
```

---

## 🔀 Pull Request Process

### Чеклист перед созданием PR

- [ ] Код следует стандартам проекта
- [ ] Запущено `go fmt ./...`
- [ ] Запущено `go vet ./...`
- [ ] Все тесты проходят
- [ ] Добавлена документация (если нужно)
- [ ] Обновлен CHANGELOG.md (если нужно)
- [ ] Ветка синхронизирована с main/develop

### Шаблон Pull Request

```markdown
## 📝 Описание

Краткое описание изменений и какую проблему они решают.

Fixes #(issue)

## 🎯 Тип изменений

- [ ] 🐛 Bug fix (исправление, не ломающее существующую функциональность)
- [ ] ✨ New feature (новая функциональность)
- [ ] 💥 Breaking change (изменения, ломающие обратную совместимость)
- [ ] 📚 Documentation (документация)
- [ ] 🎨 Style (форматирование, без изменения кода)
- [ ] ♻️ Refactoring (рефакторинг без изменения функциональности)
- [ ] ⚡ Performance (оптимизация)
- [ ] ✅ Test (добавление тестов)

## 🧪 Как тестировалось?

Опишите, как вы проверили изменения.

- [ ] Тест A
- [ ] Тест B

## 📸 Скриншоты (если применимо)

Добавьте скриншоты до/после (если UI изменения).

## ✅ Чеклист

- [ ] Код следует стандартам проекта
- [ ] Самопроверка кода выполнена
- [ ] Код прокомментирован в сложных местах
- [ ] Документация обновлена
- [ ] Изменения не генерируют новых warnings
- [ ] Тесты добавлены/обновлены
- [ ] Все тесты проходят локально
```

### Code Review Process

После создания PR:

1. ⏳ Ожидайте code review от мейнтейнеров
2. 💬 Отвечайте на комментарии
3. ♻️ Вносите запрошенные изменения
4. ✅ После одобрения PR будет смержен

### Критерии приёма

PR будет принят, если:

- ✅ Код работает и решает заявленную проблему
- ✅ Код следует стандартам проекта
- ✅ Все тесты проходят
- ✅ Нет конфликтов с main веткой
- ✅ Получен approval от минимум одного мейнтейнера

---

## 🧪 Тестирование

### Запуск тестов

```bash
# Все тесты
go test ./...

# С подробным выводом
go test -v ./...

# С coverage
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Конкретный пакет
go test ./handlers/...
```

### Написание тестов

Пример теста:

```go
// handlers/commands_test.go
package handlers

import (
    "testing"
)

func TestHandleStart(t *testing.T) {
    // Setup
    bot := setupTestBot(t)
    message := createTestMessage("/start")
    
    // Execute
    err := HandleStart(bot, message)
    
    // Assert
    if err != nil {
        t.Errorf("HandleStart returned error: %v", err)
    }
}
```

### Структура тестов

```
go_tele_bot/
├── handlers/
│   ├── commands.go
│   ├── commands_test.go      # тесты для commands.go
│   ├── messages.go
│   └── messages_test.go      # тесты для messages.go
```

---

## 📖 Документация

### Комментарии в коде

Документируйте все экспортируемые функции, типы и константы:

```go
// GetMainKeyboard возвращает главную клавиатуру для стартового экрана.
// Клавиатура содержит кнопки для основных разделов приложения.
func GetMainKeyboard() tgbotapi.InlineKeyboardMarkup {
    // ...
}
```

### README и документы

При добавлении новой функциональности:

1. Обновите **README.md** если это влияет на основную документацию
2. Создайте/обновите документы в **docs/** для детального описания
3. Добавьте примеры использования
4. Обновите **API.md** если изменяются публичные интерфейсы

### Генерация документации

```bash
# Генерация godoc
go doc ./...
go doc handlers.HandleStart

# Локальный сервер документации
godoc -http=:6060
# Откройте http://localhost:6060/pkg/go_tele_bot/
```

---

## ❓ Вопросы и помощь

### Где получить помощь?

- 📖 Проверьте [документацию](./docs/)
- 🔍 Поищите в [Issues](https://github.com/yourusername/go_tele_bot/issues)
- 💬 Создайте новый Issue с меткой `question`
- 📧 Напишите мейнтейнерам

### Полезные ресурсы

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Telegram Bot API](https://core.telegram.org/bots/api)
- [go-telegram-bot-api docs](https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api/v5)

---

## 🎉 Спасибо!

Спасибо, что вносите вклад в **AI Student Assistant**! 

Каждый Pull Request, Issue и комментарий помогают сделать проект лучше.

---

<div align="center">

**Вопросы?** Создайте [Issue](https://github.com/yourusername/go_tele_bot/issues) или свяжитесь с командой!

Made with ❤️ by the AI Student Assistant team

</div>
