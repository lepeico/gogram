package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/lepeico/gogram/pkg/telegram"
)

// TODO: make example much smaller
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
	id, err := strconv.Atoi(dev)

	gram := telegram.New(botToken)
	document, err := os.Open("LICENSE")
	if err != nil {
		panic(err)
	}
	defer document.Close()

	if js, err := gram.SendDocument(telegram.Chat{ID: id}, document); err != nil {
		log.Print(err)
	} else {
		log.Printf("%+v\n", js)
	}
}
