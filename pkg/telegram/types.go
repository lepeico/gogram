package telegram

type User struct {
	// ID is a unique identifier for this user or bot
	ID int `json:"id"`
	// IsBot true, if this user is a bot
	// Optional
	IsBot bool `json:"is_bot"`
	// FirstName user's or bot's first name
	FirstName string `json:"first_name"`
	// LastName user's or bot's last name
	// Optional
	LastName string `json:"last_name"`
	// UserName user's or bot's username
	// Optional
	UserName string `json:"username"`
	// LanguageCode IETF language tag of the user's language
	// more info: https://en.wikipedia.org/wiki/IETF_language_tag
	// Optional
	LanguageCode string `json:"language_code"`
	// CanJoinGroups true, if the bot can be invited to groups
	// returned only in getMe
	// Optional
	CanJoinGroups bool `json:"can_join_groups"`
	// CanReadAllGroupMessages true, if privacy mode is disabled for the bot
	// returned only in getMe
	// Optional
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages"`
	// SupportsInlineQueries true, if the bot supports inline queries
	// returned only in getMe
	// Optional
	SupportsInlineQueries bool `json:"supports_inline_queries"`
}
