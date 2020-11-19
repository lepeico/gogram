package gogram

import (
	"github.com/lepeico/gogram/pkg/telegram"
)

type GoGram struct {
	telegram telegram.Telegram
	BotInfo  telegram.User
}

func New(token string) *GoGram {
	g := &GoGram{}
	g.telegram = *telegram.New(token)
	return g
}

func (gg *GoGram) Launch() error {
	me, err := gg.telegram.GetMe()

	if err != nil {
		return err
	}

	gg.BotInfo = me
	return nil
}
