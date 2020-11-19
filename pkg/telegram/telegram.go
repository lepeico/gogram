package telegram

import (
	"encoding/json"
	"net/url"

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

	response, err := tg.Call("getMe", url.Values{})
	if err != nil {
		return bot, err
	}
	json.Unmarshal(response, &bot)

	return bot, nil
}

func (tg *Telegram) SendMessage(chatId string, text string) (json.RawMessage, error) {
	message := url.Values{}
	message.Set("chat_id", chatId)
	message.Set("text", text)

	response, _ := tg.Call("sendMessage", message)

	return response, nil
}
