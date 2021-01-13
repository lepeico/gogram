package telegram

func (tg *Telegram) EditMessageText(msg Message, text string, opts ...Option) (edited Message, err error) {
	if msg.Chat.ID == 0 {
		opts = append(opts, Option{"inline_message_id", msg.MessageID})
	} else {
		opts = append(opts, Option{"chat_id", msg.Chat.ID}, Option{"message_id", msg.MessageID})
	}

	err = tg.client.Call("editMessageText", createPayload(fields{
		"text": text,
	}, opts), &edited)
	return
}

//TODO
func (tg *Telegram) EditMessageCaption(msg Message, caption string, opts ...Option) (edited Message, err error) {
	if msg.Chat.ID == 0 {
		opts = append(opts, Option{"inline_message_id", msg.MessageID})
	} else {
		opts = append(opts, Option{"chat_id", msg.Chat.ID}, Option{"message_id", msg.MessageID})
	}

	err = tg.client.Call("editMessageCaption", createPayload(fields{
		"caption": caption,
	}, opts), &edited)
	return
}

func (tg *Telegram) EditMessageMedia(msg Message, media interface{}, opts ...Option) (edited Message, err error) {
	if msg.Chat.ID == 0 {
		opts = append(opts, Option{"inline_message_id", msg.MessageID})
	} else {
		opts = append(opts, Option{"chat_id", msg.Chat.ID}, Option{"message_id", msg.MessageID})
	}

	err = tg.client.Call("editMessageMedia", createPayload(fields{
		"media": media,
	}, opts), &edited)
	return
}

func (tg *Telegram) EditMessageReplyMarkup(msg Message, opts ...Option) (edited Message, err error) {
	if msg.Chat.ID == 0 {
		opts = append(opts, Option{"inline_message_id", msg.MessageID})
	} else {
		opts = append(opts, Option{"chat_id", msg.Chat.ID}, Option{"message_id", msg.MessageID})
	}

	err = tg.client.Call("editMessageReplyMarkup", createPayload(fields{}, opts), &edited)
	return
}

func (tg *Telegram) StopPoll(msg Message, opts ...Option) (poll Poll, err error) {
	err = tg.client.Call("stopPoll", createPayload(fields{
		"chat_id":    msg.Chat.ID,
		"message_id": msg.MessageID,
	}, opts), &poll)
	return
}

func (tg *Telegram) DeleteMessage(msg Message) (deleted bool, err error) {
	err = tg.client.Call("deleteMessage", createPayload(fields{
		"chat_id":    msg.Chat.ID,
		"message_id": msg.MessageID,
	}, []Option{}), &deleted)
	return
}
