package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lepeico/gogram/pkg/gogram"
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

	gram := gogram.New(botToken)

	err := gram.Launch()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", gram)
	}
}
