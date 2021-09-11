package main

import (
	"github.com/fransimanuel/RestfulApiwithHexagonalArch/app"
	"github.com/fransimanuel/RestfulApiwithHexagonalArch/logger"
	// "github.com/fransimanuel/RestfulApiwithHexagonalArch/logger"
)

func main() {
	// log.Printf("Starting Our Application")
	logger.Info("Starting Our Application")
	app.Start()
}
