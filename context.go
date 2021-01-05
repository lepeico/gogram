package gogram

import "github.com/lepeico/gogram/pkg/telegram"

type Context struct {
	*telegram.Update

	Telegram *telegram.Telegram
	Bot      *telegram.User
	State    map[string]interface{}
}

func (c *Context) Reply(text string, opts ...telegram.Option) (msg telegram.Message, err error) {
	return c.Telegram.SendMessage(*c.Update.Message.Chat, text, opts...)
}
