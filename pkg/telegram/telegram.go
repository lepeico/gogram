package telegram

import (
	"encoding/json"
	"net/url"

	"github.com/lepeico/gogram/internal/client"
)

type Telegram struct {
	httpClient *client.Client
}

func New(token string) *Telegram {
	return &Telegram{httpClient: client.NewClient(token)}
}

func (t *Telegram) GetMe() (User, error) {
	var bot User

	response, err := t.httpClient.Call("getMe", url.Values{})
	if err != nil {
		return bot, err
	}
	json.Unmarshal(response, &bot)

	return bot, nil
}
