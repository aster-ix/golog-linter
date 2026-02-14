# Golog-linter

Кастомный линтер для проверки логов в Go коде. Интегрируется с golangci-lint v2.9.0+ через module plugin system

## Какие правила учтены

1. **Логи должны начинаться с маленькой буквы**
   ```go
   log.Println("Server started") // ошибка
   log.Println("server started") // правильно
   ```
2. **Логи только на английском**
   ```go
   log.Println("запуск сервера") // ошибка
   log.Println("server started") // правильно
   ```
3. **Логи без спецсимволов**
   ```go
   log.Println("error!!!") // ошибка
   log.Println("error")    // правильно
   ```
4. **Логи без переменных (только строковые литералы)**
   ```go
   var msg string
   log.Println(msg)              // ошибка
   log.Println("server started") // правильно
   ```

## Поддерживаемые библиотеки

* `log` (стандартная библиотека)
* `log/slog` (стандартная библиотека)
* `go.uber.org/zap`

## Требования

* Go 1.25.0+
* golangci-lint v2.9.0+

## Установка и сборка

### 1. Клонируйте репозиторий

```bash
git clone https://github.com/aster-ix/golog-linter.git
cd golog-linter
```

### 2. Установите зависимости

```bash
go mod tidy
```

### 3. Соберите кастомный golangci-lint

```bash
golangci-lint custom -v
```

Это создаст исполняемый файл `gologlinter.exe` (Windows) или `gologlinter` (Linux/macOS).

## Использование

### Базовое использование

```bash
# Проверить весь проект
./gologlinter.exe run

# Проверить конкретную директорию
./gologlinter.exe run ./pkg/...

# Проверить конкретный файл
./gologlinter.exe run ./main.go
```

## Структура проекта

```
golog-linter/
├── analyzer/
│   ├── analyzer_test.go
│   └── analyzer.go       # Основная логика проверок
├── testdata/
│   ├── go.sum
│   ├── tests.go          # Тестовые данные
│   └── go.mod   
├── plugin.go             # Интеграция с golangci-lint
├── go.mod  
├── go.sum
├── .custom-gcl.yml       # Конфигурация сборки
├── .golangci.yml         # Конфигурация линтера
├── makefile              # Команды для сборки
└── README.md
```

## Makefile команды

```bash
# Показать все доступные команды
make help

# Собрать кастомный golangci-lint (основная команда)
make plugin

# Собрать executable
make exec

# Запустить тесты (через go test)
make test

# Собрать и запустить на тестовых данных
make run

# Очистить артефакты сборки
make clean

# Собрать всё (по умолчанию = plugin)
make
```
## Автор

[@aster-ix](https://github.com/aster-ix)
