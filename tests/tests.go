package tests

import (
	"log"
	"log/slog"
)

func tests() {
	log.Println("Bad log1")
	slog.Info("Bad log2")
}
