package telegram

// Game represents a game. 
// Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
	// Title of the game
	Title string `json:"title"`
	// Description of the game
	Description string `json:"description"`
	// Photo that will be displayed in the game message in chats.
	Photo *[]PhotoSize `json:"photo"`
	// Text a brief description of the game or high scores included in the game message.
	// Can be automatically edited to include current high scores for the game
	// when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
	// Optional
	Text string `json:"text"`
	// TextEntities special entities that appear in text, such as usernames, URLs, bot commands, etc.
	// Optional
	TextEntities *[]MessageEntity `json:"text_entities"`
	// Animation animation that will be displayed in the game message in chats.
	// Upload via BotFather (https://t.me/botfather).
	// Optional
	Animation *Animation `json:"animation"`
}

// CallbackGame a placeholder, currently holds no information. 
// Use BotFather to set up your game.
type CallbackGame struct{}

// GameHighScore is a user's score and position on the leaderboard.
type GameHighScore struct {
	// Position in high score table for the game
	Position int `json:"position"`
	// User user
	User *User `json:"user"`
	// Score score
	Score int `json:"score"`
}