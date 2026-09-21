# Changelog

Все значимые изменения в проекте **AI Student Assistant** будут документироваться в этом файле.

Формат основан на [Keep a Changelog](https://keepachangelog.com/ru/1.0.0/),
и проект следует [Semantic Versioning](https://semver.org/lang/ru/).

---

## [Unreleased]

### Planned
- 📅 Система расписания
- 🗄️ Интеграция PostgreSQL
- 🔔 Push-уведомления
- 🏫 Информация о кабинетах
- 👨‍🏫 Данные о преподавателях

---

## [0.1.0] - 2026-09-19

### Added
- ✨ Начальная инициализация проекта
- ✨ Настройка Telegram Bot API
- ✨ Polling система для получения updates
- ✨ Базовая архитектура handlers
- ✨ Структура для commands, messages, callbacks
- ✨ Модуль keyboards для inline клавиатур
- ✨ Конфигурация через .env
- 📝 Красивый README.md с документацией
- 📝 CONTRIBUTING.md с руководством для разработчиков
- 📝 CODE_OF_CONDUCT.md
- 📝 LICENSE (MIT)
- 🔧 .gitignore с правилами для Go, IDE, OS
- 🔧 .env.example с примером конфигурации

### Project Structure
```
go_tele_bot/
├── bot/settings.go       # Конфигурация бота
├── handlers/             # Обработчики событий
├── keyboards/            # Генерация клавиатур
├── main.go               # Точка входа
└── docs/                 # Документация
```

### Dependencies
- go-telegram-bot-api v5.5.1
- godotenv v1.5.1

---

## Типы изменений

- `Added` - новая функциональность
- `Changed` - изменения в существующей функциональности
- `Deprecated` - функциональность, которая скоро будет удалена
- `Removed` - удалённая функциональность
- `Fixed` - исправления ошибок
- `Security` - исправления безопасности

---

## Версионирование

Проект использует [Semantic Versioning](https://semver.org/):

- **MAJOR** версия - несовместимые изменения API
- **MINOR** версия - обратно совместимая новая функциональность
- **PATCH** версия - обратно совместимые исправления ошибок

Пример: `1.2.3` где:
- `1` - MAJOR
- `2` - MINOR
- `3` - PATCH

---

<div align="center">

[Unreleased]: https://github.com/yourusername/go_tele_bot/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/yourusername/go_tele_bot/releases/tag/v0.1.0

</div>
