package telegram

import (
	"encoding/json"

	"github.com/lepeico/gogram/internal/client"
)

type Telegram struct {
	client.Client
}

func New(token string) *Telegram {
	return &Telegram{*client.NewClient(token)}
}

func (tg *Telegram) GetMe() (User, error) {
	var bot User

	response, err := tg.Call("getMe", nil)
	if err != nil {
		return bot, err
	}
	json.Unmarshal(response, &bot)

	return bot, nil
}

func (tg *Telegram) SendMessage(chatID string, text string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("text", text)

	res, err = tg.Call("sendMessage", message)
	return
}

func (tg *Telegram) SendDocument(chatID string, document interface{}) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("document", document)

	res, err = tg.Call("sendDocument", message)
	return
}
