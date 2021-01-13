package telegram

type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	PollAnswer         *PollAnswer         `json:"poll_answer,omitempty"`
}

var AllowedUpdates []string = []string{
	"message",
	"edited_message",
	"channel_post",
	"edited_channel_post",
	"callback_query",
	"inline_query",
	"chosen_inline_result",
	"shipping_query",
	"pre_checkout_query",
	"poll",
	"poll_answer",
}

func (tg *Telegram) GetUpdates(offset, limit, timeout int, allowedUpdates []string, opts ...Option) (upd []Update, err error) {
	err = tg.client.Call("getUpdates", createPayload(fields{
		"offset":          offset,
		"limit":           limit,
		"timeout":         timeout,
		"allowed_updates": allowedUpdates,
	}, opts), &upd)
	return
}
