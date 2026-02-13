package tests

import (
	"fmt"
	"log"
	"log/slog"
)

func tests() {
	log.Println("Bad log1")
	slog.Info("Bad log2")

	fmt.Println("not log")
}
