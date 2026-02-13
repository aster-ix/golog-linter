package tests

import (
	"fmt"
	"log"
	"log/slog"
)

func tests() {
	log.Println("Bad log1")
	slog.Info("Bad log2")

	log.Println("good log1")
	slog.Info("gooD log2")

	log.Println("Log in рussian")

	fmt.Println("not log")
}
