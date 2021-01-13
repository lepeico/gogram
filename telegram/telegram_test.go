package telegram

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var gram *Telegram
var devID int

var envPath = "./../../.env"
var filePath = "./../../LICENSE"
var mediaPath = "./../../examples/telegram/assets/test.mp4"

func gramInit() {
	botToken, exists := os.LookupEnv("TELEGRAM_TOKEN")
	if !exists {
		os.Exit(2)
	}

	id, exists := os.LookupEnv("DEV_ID")
	if !exists {
		os.Exit(2)
	}

	devID, _ = strconv.Atoi(id)

	//New Telegram
	gram = New(botToken)
}

func TestMain(m *testing.M) {
	// loads values from .env into the system
	if err := godotenv.Load(envPath); err != nil {
		log.Print("No .env file found")
	}
	gramInit()
	os.Exit(m.Run())
}

//TODO:
// func TestSendMessage() {

// }

//TODO:
func TestSendDocument(t *testing.T) {
	document, err := os.Open(filePath)
	if err != nil {
		t.Error()
	}
	defer document.Close()

	if _, err := gram.SendDocument(Chat{ID: devID}, document); err != nil {
		t.Error(err)
	}
}
