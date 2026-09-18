🎓 Go Telegram Bot

<p align="center">
  <strong>Telegram-бот для автоматизации образовательного процесса</strong>
</p>
<p align="center">
  Go • PostgreSQL • Telegram
</p>

⸻

📌 О проекте

Go Telegram Bot — проект, направленный на создание единого цифрового помощника для студентов.

Основная задача проекта — собрать часто используемые образовательные функции в одном Telegram-боте и сделать взаимодействие с учебной информацией быстрым и удобным.

Проект разрабатывается с упором на модульную архитектуру, чтобы в дальнейшем его можно было расширять без переработки основной системы.

⸻

🎯 Цели проекта

На текущем этапе проект развивается вокруг нескольких основных направлений:

* 📅 получение расписания занятий;
* 📰 просмотр новостей;
* 🏫 навигация по учебному корпусу;
* 🔄 отображение изменений расписания;
* 📚 работа с учебными материалами;
* 🤖 интеллектуальный помощник;
* 🗄️ хранение и обработка данных;
* 🔌 возможность дальнейшего расширения на другие мессенджеры.

⸻

🚧 Текущий статус

In Development

Сейчас разрабатывается базовая архитектура Telegram-бота и инфраструктура проекта.

На первом этапе реализуется:

Telegram
   ↓
Bot API
   ↓
Polling
   ↓
Handlers
   ↓
Commands / Messages / Callbacks
   ↓
Keyboards

После формирования базовой архитектуры начнётся разработка основной функциональности.

⸻

🔮 Будущее проекта

Проект планируется развивать поэтапно.

Этап 1 — Core

* Инициализация Go-проекта
* Подключение Telegram Bot API
* Настройка polling
* Конфигурация через environment variables
* Обработка команд
* Обработка сообщений
* Inline-клавиатуры
* Callback handlers

Этап 2 — Education

* Расписание
* Выбор группы
* Расписание по дням
* Новости
* Навигация по кабинетам
* Изменения расписания

Этап 3 — Database

* PostgreSQL
* Модели данных
* Хранение групп и расписаний
* Работа с пользователями

Этап 4 — Intelligent Features

* Работа с учебными материалами
* Поиск по материалам
* Краткие конспекты
* Ответы на вопросы по учебным материалам

Этап 5 — Multi-Messenger

В перспективе архитектура должна позволять использовать одну backend-систему для нескольких платформ:

                 ┌──────────────┐
                 │    Backend   │
                 └──────┬───────┘
                        │
          ┌─────────────┼─────────────┐
          ↓             ↓             ↓
      Telegram         MAX         Other

⸻

🏗️ Архитектура

go_tele_bot/
│
├── bot/
│   └── settings.go
│
├── handlers/
│   ├── handler.go
│   ├── commands.go
│   ├── messages.go
│   └── callbacks.go
│
├── keyboards/
│   └── keyboard.go
│
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md

Ответственность компонентов

Компонент	Назначение
main.go	Точка входа и запуск приложения
bot/settings.go	Конфигурация Telegram-бота
handlers/handler.go	Маршрутизация Telegram updates
handlers/commands.go	Обработка команд
handlers/messages.go	Обработка текстовых сообщений
handlers/callbacks.go	Обработка callback-запросов
keyboards/keyboard.go	Создание клавиатур
.env	Локальные секреты и переменные окружения

⸻

🛠️ Tech Stack

Technology	Purpose
Go	Backend
Telegram Bot API	Взаимодействие с Telegram
PostgreSQL	Хранение данных
godotenv	Загрузка .env
Git	Контроль версий
GitHub	Совместная разработка

⸻

⚙️ Запуск проекта

1. Клонирование

git clone <repository-url>
cd go_tele_bot

2. Установка зависимостей

go mod download

3. Настройка .env

Создайте файл .env в корне проекта:

TELEGRAM_BOT_TOKEN=your_token

⚠️ .env не должен попадать в Git.

Для удобства в репозитории используется .env.example.

4. Запуск

go run main.go

⸻

🌿 Git Workflow

Для разработки используются отдельные ветки.

main
│
├── feature/...
├── fix/...
├── refactor/...
└── docs/...

Названия веток

feature/add-schedule
feature/inline-keyboard
fix/callback-handler
refactor/handlers
docs/update-readme

Не рекомендуется выполнять разработку непосредственно в main.

⸻

📝 Commit Convention

Для сообщений коммитов используется подход Conventional Commits.

Формат:

<type>: <description>

Типы коммитов

Type	Назначение
feat	Новая функциональность
fix	Исправление ошибки
refactor	Изменение структуры кода без изменения поведения
docs	Документация
test	Тесты
chore	Технические изменения
style	Форматирование и стиль
perf	Оптимизация производительности

Примеры

feat: add schedule handler
feat: add inline keyboard
fix: handle empty message
refactor: separate telegram handlers
docs: update project readme
test: add schedule tests
chore: configure environment

Правила

1. Один коммит — одно логическое изменение.
2. Сообщение должно кратко описывать изменение.
3. Используется английский язык.
4. Не использовать сообщения вроде:

fix
update
changes
new
test
asdf
final

5. Не смешивать в одном коммите несвязанные изменения.

⸻

👨‍💻 Developers

<table>
  <tr>
    <td align="center">
      <strong>devskyyy</strong><br>
      Developer
    </td>
    <td align="center">
      <strong>pupsmane</strong><br>
      Developer
    </td>
  </tr>
</table>

⸻

📄 License

The project is currently being developed for educational and research purposes.

⸻

<p align="center">
  <strong>Go Telegram Bot</strong><br>
  Built with Go 🚀
</p>