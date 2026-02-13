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

	log.Println("log in russian23!!")

	fmt.Println("not log")
}
