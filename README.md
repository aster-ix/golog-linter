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
go mod tidy -C testdata
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
# Показать список всех доступных команд
make help

# Собрать кастомный golangci-lint с плагином (основная команда)
make plugin

# Собрать standalone executable (gologlinter)
make exec

# Запустить unit-тесты (go test ./analyzer)
make test

# Собрать плагин и запустить линтер на testdata
make run

# Запустить линтер напрямую на testdata
make test_w_linter

# Очистить артефакты сборки
make clean

# Собрать (по умолчанию выполняется make plugin)
make

```
## Автор

[@aster-ix](https://github.com/aster-ix)
