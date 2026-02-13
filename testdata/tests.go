package a

import (
	"log"
	"log/slog"
)

func test1() {
	log.Println("Big") // want "log should start with lower case"
	slog.Info("Start") // want "log should start with lower case"

	log.Println("small") // OK
	slog.Info("good")    // OK
}

func test2() {
	log.Println("привет") // want "log should be only in English"
	slog.Info("ошибка")   // want "log should be only in English"

	log.Println("hello") // OK
	slog.Info("error")   // OK
}

func test3() {
	log.Println("bad!")  // want "log should not contain symbols"
	slog.Info("wait...") // want "log should not contain symbols"
	log.Printf("ok🚀")    // want "log should not contain symbols"

	log.Println("good") // OK
	slog.Info("done")   // OK
}

func test4() {
	var msg string

	log.Println(msg)         // want "log should not contain variables for safety"
	log.Printf("text" + msg) // want "log should not contain variables for safety"

	log.Println("static") // OK
	slog.Info("text")     // OK
}

func test5() {
	log.Println("Bad!")  // want "log should start with lower case" "log should not contain symbols"
	slog.Info("привет!") // want "log should be only in English" "log should not contain symbols"

	var v string
	log.Println(v) // want "log should not contain variables for safety"
}
