package telegram

type Option struct {
	Name      string
	Parameter interface{}
}

func OptionParseMode(parseMode string) Option {
	return Option{Name: "parse_mode", Parameter: parseMode}
}

func OptionParseModeHTML() Option {
	return OptionParseMode("HTML")
}

func OptionParseModeMarkdown() Option {
	return OptionParseMode("Markdown")
}

func OptionParseModeMarkdownV2() Option {
	return OptionParseMode("MarkdownV2")
}

func OptionDisableNotification() Option {
	return Option{Name: "disable_notification", Parameter: true}
}

func OptionDisableWebPagePreview() Option {
	return Option{Name: "disable_web_page_preview", Parameter: true}
}

func OptionReplyToMessageId(messageID int) Option {
	return Option{Name: "reply_to_message_id", Parameter: messageID}
}

// Default to '🎲'
func OptionEmoji(emoji string) Option {
	return Option{Name: "emoji", Parameter: emoji}
}

func OptionLimit(limit int) Option {
	return Option{Name: "limit", Parameter: limit}
}
