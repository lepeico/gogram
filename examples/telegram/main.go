package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lepeico/gogram/pkg/telegram"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	botToken, exists := os.LookupEnv("TELEGRAM_TOKEN")
	if !exists {
		os.Exit(2)
	}

	dev, exists := os.LookupEnv("DEV_ID")
	if !exists {
		os.Exit(2)
	}

	gram := telegram.New(botToken)
	document, err := os.Open("LICENSE")
	if err != nil {
		panic(err)
	}
	defer document.Close()

	if js, err := gram.SendDocument(dev, document); err != nil {
		log.Print(err)
	} else {
		log.Printf("%+v\n", string(js))
	}
}
