package a

import (
	"fmt"
	"log"
	"log/slog"

	"go.uber.org/zap"
)

func testRule1() {
	log.Println("Big")    // want "log should start with lower case"
	slog.Info("Start")    // want "log should start with lower case"
	log.Printf("Warning") // want "log should start with lower case"
	slog.Error("Failed")  // want "log should start with lower case"

	log.Println("small")
	slog.Info("good")
	log.Printf("starting")
}

func testRule2() {
	log.Println("запуск сервера")        // want "log should be only in English"
	slog.Info("ошибка подключения")      // want "log should be only in English"
	log.Printf("сервер запущен")         // want "log should be only in English"
	slog.Error("база данных недоступна") // want "log should be only in English"

	log.Println("server started")
	slog.Info("connection established")
	log.Printf("processing request")
}

func testRule3() {
	log.Println("server started!")    // want "log should not contain symbols"
	slog.Info("connection failed!!!") // want "log should not contain symbols"
	log.Printf("waiting...")          // want "log should not contain symbols"
	log.Println("server started🚀")    // want "log should not contain symbols"

	log.Println("server started")
	slog.Info("connection established")
	log.Printf("processing file")
}

func testRule4() {

	var msg string

	log.Println(msg)                     // want "log should not contain variables for safety"
	log.Println("user password: " + msg) // want "log should not contain variables for safety"
	slog.Info("api key: " + msg)         // want "log should not contain variables for safety"
	slog.Error("token: " + msg)          // want "log should not contain variables for safety"

	log.Println("user authenticated")
	slog.Info("api request completed")
	slog.Error("token validated")
}

func testCombined() {
	log.Println("Bad")   // want "log should start with lower case"
	slog.Info("Привет")  // want "log should be only in English"
	log.Printf("error!") // want "log should not contain symbols"

	var v string
	log.Println(v) // want "log should not contain variables for safety"

	log.Println("good message")
	slog.Info("everything ok")
}

func testEdgeCases() {
	log.Println("")                    // OK - пустая строка
	log.Println("   ")                 // OK - только пробелы
	log.Println("12345")               // OK - только цифры
	log.Println("server 8080 started") // OK - цифры с текстом
	log.Println("a")                   // OK - одна маленькая буква
	log.Println("A")                   // want "log should start with lower case"
}

func testNonLogFunctions() {
	fmt.Println("Test")
	println("Starting Server!")
	print("Error occurred!")
}

func testZapRule1() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	logger.Info("Starting server")    // want "log should start with lower case"
	logger.Error("Failed to connect") // want "log should start with lower case"
	sugar.Infow("Server started")     // want "log should start with lower case"

	logger.Info("starting server")
	sugar.Infow("server started")
}
