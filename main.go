package main

import (
	"io"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("app.log", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("")
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(logFile, os.Stdout)
	logger := log.New(multiWriter, "Info: ", log.LstdFlags)

	// Log first execution.
	logger.Println("Program has been executed..")
}
