package gogram

import t "github.com/lepeico/gogram/telegram"

type Context struct {
	*t.Update

	Telegram *t.Telegram
	Bot      *t.User
	State    map[string]interface{}
}

func (u *Context) Type() (updateEntitiesTypes []string) {
	var entities *[]t.MessageEntity
	switch {
	case u.Message != nil:
		entities = u.Message.Entities
	case u.EditedMessage != nil:
		entities = u.EditedMessage.Entities
	}

	for i, entity := range *entities {
		updateEntitiesTypes[i] = entity.Type
	}

	return
}

func (u *Context) From() (from *t.User) {
	switch {
	case u.Message != nil:
		from = u.Message.From
	case u.EditedMessage != nil:
		from = u.EditedMessage.From
	case u.CallbackQuery != nil:
		from = u.CallbackQuery.From
	case u.InlineQuery != nil:
		from = u.InlineQuery.From
	case u.ChannelPost != nil:
		from = u.ChannelPost.From
	case u.EditedChannelPost != nil:
		from = u.EditedChannelPost.From
	case u.ShippingQuery != nil:
		from = u.ShippingQuery.From
	case u.PreCheckoutQuery != nil:
		from = u.PreCheckoutQuery.From
	case u.ChosenInlineResult != nil:
		from = u.ChosenInlineResult.From
	}

	return
}

func (u *Context) Chat() (chat *t.Chat) {
	switch {
	case u.Message != nil:
		chat = u.Message.Chat
	case u.EditedMessage != nil:
		chat = u.EditedMessage.Chat
	case u.CallbackQuery != nil:
		chat = u.CallbackQuery.Message.Chat
	case u.ChannelPost != nil:
		chat = u.ChannelPost.Chat
	case u.EditedChannelPost != nil:
		chat = u.EditedChannelPost.Chat
	}

	return
}

func (c *Context) Reply(text string, opts ...t.Option) (msg t.Message, err error) {
	return c.Telegram.SendMessage(*c.Chat(), text, opts...)
}
