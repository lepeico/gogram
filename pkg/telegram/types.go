package telegram

// User contains information about user
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

// Chat contains information about the place a message was sent.
type Chat struct {
	// ID is a unique identifier for this chat
	ID int `json:"id"`
	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`
	// Title for supergroups, channels and group chats
	// Optional
	Title string `json:"title"`
	// UserName for private chats, supergroups and channels if available
	// Optional
	UserName string `json:"username"`
	// FirstName of the other party in a private chat
	// Optional
	FirstName string `json:"first_name"`
	// LastName of the other party in a private chat
	// Optional
	LastName string `json:"last_name"`
	// Photo is a chat photo
	// Optional
	Photo *ChatPhoto `json:"photo"`
	// Bio of the other party in a private chat
	// Returned only in getChat
	// Optional
	Bio string `json:"bio"`
	// Description for groups, supergroups and channel chats
	// Optional
	Description string `json:"description"`
	// InviteLink is a chat invite link, for groups, supergroups and channel chats.
	// Each administrator in a chat generates their own invite links,
	// so the bot must first generate the link using exportChatInviteLink
	// Optional
	InviteLink string `json:"invite_link"`
	// PinnedMessage Pinned message, for groups, supergroups and channels
	// Optional
	PinnedMessage *Message `json:"pinned_message"`
	// Default chat member permissions, for groups and supergroups.
	// Returned only in getChat
	// Optional
	Permissions *ChatPermissions `json:"permissions"`
	// For supergroups, the minimum allowed delay between consecutive
	// messages sent by each unpriviledged user.
	// Returned only in getChat
	// Optional
	SlowModeDelay int `json:"slow_mode_delay"`
	// For supergroups, name of group sticker set.
	// Returned only in getChat
	// Optional
	StickerSetName string `json:"sticker_set_name"`
	// True, if the bot can change the group sticker set.
	// Returned only in getChat
	// Optional
	CanSetStickerSet bool `json:"can_set_sticker_set"`
	// Unique identifier for the linked chat, i.e. the discussion group
	// identifier for a channel and vice versa; for supergroups and channel chats.
	// Returned only in getChat
	// Optional
	LinkedChatID int `json:"linked_chat_id"`
	// For supergroups, the location to which the supergroup is connected.
	// Returned only in getChat
	// Oprional
	Location *ChatLocation `json:"location"`
}

// Message is returned by almost every request, and contains data about
// almost anything.
type Message struct {
	// MessageID is a unique message identifier inside this chat
	MessageID int `json:"message_id"`
	// From is a sender, empty for messages sent to channels;
	// Optional
	From *User `json:"from"`
	// Sender of the message, sent on behalf of a chat.
	//The channel itself for channel messages.
	//The supergroup itself for messages from anonymous group administrators.
	//The linked channel for messages automatically forwarded to the discussion group
	// Optional
	SenderChat *Chat `json:"sender_chat"`
	// Date of the message was sent in Unix time
	Date int `json:"date"`
	// Chat is the conversation the message belongs to
	Chat *Chat `json:"chat"`
	// ForwardFrom for forwarded messages, sender of the original message;
	// Optional
	ForwardFrom *User `json:"forward_from"`
	// ForwardFromChat for messages forwarded from channels,
	// information about the original channel;
	// Optional
	ForwardFromChat *Chat `json:"forward_from_chat"`
	// ForwardFromMessageID for messages forwarded from channels,
	// identifier of the original message in the channel;
	// Optional
	ForwardFromMessageID int `json:"forward_from_message_id"`
	// For messages forwarded from channels, signature of the post author if present
	// Optional
	ForwardSignature string `json:"forward_signature"`
	// Sender's name for messages forwarded from users who disallow adding
	// a link to their account in forwarded messages
	ForwardSenderName string `json:"forward_sender_name"`
	// ForwardDate for forwarded messages, date the original message was sent in Unix time;
	// Optional
	ForwardDate int `json:"forward_date"`
	// ReplyToMessage for replies, the original message.
	// Note that the Message object in this field will not contain further ReplyToMessage fields
	// even if it itself is a reply
	// Optional
	ReplyToMessage *Message `json:"reply_to_message"`
	// ViaBot is Bot through which the message was sent;
	// Optional
	ViaBot *User `json:"via_bot"`
	// EditDate of the message was last edited in Unix time;
	// Optional
	EditDate int `json:"edit_date"`
	// MediaGroupID is the unique identifier of a media message group this message belongs to;
	// Optional
	MediaGroupID string `json:"media_group_id"`
	// AuthorSignature is the signature of the post author for messages in channels;
	// Optional
	AuthorSignature string `json:"author_signature"`
	// Text is for text messages, the actual UTF-8 text of the message, 0-4096 characters;
	// optional
	Text string `json:"text"`
	// Entities is for text messages, special entities like usernames,
	// URLs, bot commands, etc. that appear in the text;
	// Optional
	Entities *[]MessageEntity `json:"entities"`
	// Message is an animation, information about the animation.
	// For backward compatibility, when this field is set,
	// the document field will also be set
	// Optional
	Animation *Animation `json:"animation"`
	// Audio message is an audio file, information about the file;
	// Optional
	Audio *Audio `json:"audio"`
	// Document message is a general file, information about the file;
	// Optional
	Document *Document `json:"document"`
	// Message is a photo, available sizes of the photo;
	// Optional
	Photo *[]PhotoSize `json:"photo"`
	// Message is a sticker, information about the sticker;
	// Optional
	Sticker *Sticker `json:"sticker"`
	// Message is a video, information about the video;
	// Optional
	Video *Video `json:"video"`
	// Message is a video note, information about the video message;
	// Optional
	VideoNote *VideoNote `json:"video_note"`
	// Message is a voice message, information about the file;
	// Optional
	Voice *Voice `json:"voice"`
	// Caption for the animation, audio, document, photo, video or voice, 0-1024 characters;
	// Optional
	Caption string `json:"caption"`
	// CaptionEntities for messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// Message is a shared contact, information about the contact;
	// Optional
	Contact *Contact `json:"contact"`
	// Message is a dice with random value from 1 to 6
	// Optional
	Dice *Dice `json:"dice"`
	// Message is a game, information about the game;
	// Optional
	Game *Game `json:"game"`
	// Message is a native poll, information about the poll
	// Optional
	Poll *Poll `json:"poll"`
	// Message is a venue, information about the venue.
	// For backward compatibility, when this field is set, the location field will also be set;
	// Optional
	Venue *Venue `json:"venue"`
	// Message is a shared location, information about the location;
	// Optional
	Location *Location `json:"location"`
	// New members that were added to the group or supergroup
	// and information about them (the bot itself may be one of these members);
	// Optional
	NewChatMembers *[]User `json:"new_chat_members"`
	// LeftChatMember is a member was removed from the group,
	// information about them (this member may be the bot itself);
	// Optional
	LeftChatMember *User `json:"left_chat_member"`
	// NewChatTitle is a chat title was changed to this value;
	// Optional
	NewChatTitle string `json:"new_chat_title"`
	// NewChatPhoto is a chat photo was change to this value;
	// Optional
	NewChatPhoto *[]PhotoSize `json:"new_chat_photo"`
	// DeleteChatPhoto is a service message: the chat photo was deleted;
	// Optional
	DeleteChatPhoto bool `json:"delete_chat_photo"`
	// GroupChatCreated is a service message: the group has been created;
	// Optional
	GroupChatCreated bool `json:"group_chat_created"`
	// SuperGroupChatCreated is a service message: the supergroup has been created.
	// This field can't be received in a message coming through updates,
	// because bot can't be a member of a supergroup when it is created.
	// It can only be found in ReplyToMessage if someone replies to a very first message
	// in a directly created supergroup;
	// Optional
	SuperGroupChatCreated bool `json:"supergroup_chat_created"`
	// ChannelChatCreated is a service message: the channel has been created.
	// This field can't be received in a message coming through updates,
	// because bot can't be a member of a channel when it is created.
	// It can only be found in ReplyToMessage
	// if someone replies to a very first message in a channel;
	// Optional
	ChannelChatCreated bool `json:"channel_chat_created"`
	// MigrateToChatID is the group has been migrated to a supergroup with the specified identifier.
	// This number may be greater than 32 bits and some programming languages
	// may have difficulty/silent defects in interpreting it.
	// But it is smaller than 52 bits, so a signed 64 bit integer
	// or double-precision float type are safe for storing this identifier;
	// Optional
	MigrateToChatID int `json:"migrate_to_chat_id"`
	// MigrateFromChatID is the supergroup has been migrated from a group with the specified identifier.
	// This number may be greater than 32 bits and some programming languages
	// may have difficulty/silent defects in interpreting it.
	// But it is smaller than 52 bits, so a signed 64 bit integer
	// or double-precision float type are safe for storing this identifier;
	//
	// optional
	MigrateFromChatID int `json:"migrate_from_chat_id"`
	// PinnedMessage is a specified message was pinned.
	// Note that the Message object in this field will not contain further ReplyToMessage
	// fields even if it is itself a reply;
	// Optional
	PinnedMessage *Message `json:"pinned_message"`
	// Invoice message is an invoice for a payment;
	// Optional
	Invoice *Invoice `json:"invoice"`
	// SuccessfulPayment message is a service message about a successful payment,
	// information about the payment;
	// Optional
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment"`
	// The domain name of the website on which the user has logged in
	// Optional
	ConnectedWebsite string `json:"connected_website"`
	// PassportData is a Telegram Passport data;
	// Optional
	PassportData *PassportData `json:"passport_data"`
	// ProximityAlertTriggered is a service message.
	//A user in the chat triggered another user's proximity alert while sharing Live Location
	//Optional
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered"`
	// ReplyMarkup is a inline keyboard attached to the message.
	// login_url buttons are represented as ordinary url buttons.
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}

// MessageEntity contains information about data in a Message.
type MessageEntity struct {
	// Type of the entity.
	// Can be:
	//  “mention” (@username),
	//  “hashtag” (#hashtag),
	//  “cashtag” ($USD),
	//  “bot_command” (/start@jobs_bot),
	//  “url” (https://telegram.org),
	//  “email” (do-not-reply@telegram.org),
	//  “phone_number” (+1-212-555-0123),
	//  “bold” (bold text),
	//  “italic” (italic text),
	//  “underline” (underlined text),
	//  “strikethrough” (strikethrough text),
	//  “code” (monowidth string),
	//  “pre” (monowidth block),
	//  “text_link” (for clickable text URLs),
	//  “text_mention” (for users without usernames)
	Type string `json:"type"`
	// Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`
	// Length of the entity in UTF-16 code units
	Length int `json:"length"`
	// URL for “text_link” only, url that will be opened after user taps on the text
	// Optional
	URL string `json:"url"`
	// User for “text_mention” only, the mentioned user
	// Optional
	User *User `json:"user"`
	// For “pre” only, the programming language of the entity text
	// Optional
	Language string `json:"language"`
}

// PhotoSize contains information about photos.
type PhotoSize struct {
	// FileID identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Width photo width
	Width int `json:"width"`
	// Height photo height
	Height int `json:"height"`
	// FileSize file size
	// Optional
	FileSize int `json:"file_size"`
}

// Animation is a GIF animation demonstrating the game.
type Animation struct {
	// FileID identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Video width as defined by sender
	Width int `json:"width"`
	// Video height as defined by sender
	Height int `json:"height"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Thumb animation thumbnail as defined by sender.
	// Optional
	Thumb *PhotoSize `json:"thumb"`
	// FileName original animation filename as defined by sender.
	// Optional
	FileName string `json:"file_name"`
	// MimeType of the file as defined by sender.
	// Optional
	MimeType string `json:"mime_type"`
	// FileSize ile size
	// Optional
	FileSize int `json:"file_size"`
}

// Audio contains information about audio.
type Audio struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`
	// Performer of the audio as defined by sender or by audio tags
	// Optional
	Performer string `json:"performer"`
	// Title of the audio as defined by sender or by audio tags
	// Optional
	Title string `json:"title"`
	// FileName original  filename as defined by sender.
	// Optional
	FileName string `json:"file_name"`
	// MimeType of the file as defined by sender
	// Optional
	MimeType string `json:"mime_type"`
	// FileSize file size
	// Optional
	FileSize int `json:"file_size"`
	// Thumbnail of the album cover to which the music file belongs
	// Optional
	Thumb *PhotoSize `json:"thumb"`
}

// Document contains information about a document.
type Document struct {
	// FileID is a identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Document thumbnail as defined by sender
	// Optional
	Thumb *PhotoSize `json:"thumb"`
	// FileName original filename as defined by sender
	// Optional
	FileName string `json:"file_name"`
	// MimeType  of the file as defined by sender
	// Optional
	MimeType string `json:"mime_type"`
	// FileSize file size
	// Optional
	FileSize int `json:"file_size"`
}

// Video contains information about a video.
type Video struct {
	// FileID identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Width video width as defined by sender
	Width int `json:"width"`
	// Height video height as defined by sender
	Height int `json:"height"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Thumb video thumbnail
	// Optional
	Thumb *PhotoSize `json:"thumb"`
	// FileName original filename as defined by sender
	// Optional
	FileName string `json:"file_name"`
	// MimeType of a file as defined by sender
	// Optional
	MimeType string `json:"mime_type"`
	// FileSize file size
	// Optional
	FileSize int `json:"file_size"`
}

// VideoNote contains information about a video.
type VideoNote struct {
	// FileID identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Length video width and height (diameter of the video message) as defined by sender
	Length int `json:"length"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Thumb video thumbnail
	// Optional
	Thumb *PhotoSize `json:"thumb"`
	// FileSize file size
	// Optional
	FileSize int `json:"file_size"`
}

// Voice contains information about a voice.
type Voice struct {
	// FileID identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`
	// MimeType of the file as defined by sender
	// Optional
	MimeType string `json:"mime_type"`
	// FileSize file size
	// Optional
	FileSize int `json:"file_size"`
}

// Contact contains information about a contact.
type Contact struct {
	// PhoneNumber contact's phone number
	PhoneNumber string `json:"phone_number"`
	// FirstName contact's first name
	FirstName string `json:"first_name"`
	// LastName contact's last name
	// Optional
	LastName string `json:"last_name"`
	// UserID contact's user identifier in Telegram
	// Optional
	UserID int `json:"user_id"`
	// Additional data about the contact in the form of a vCard
	Vcard string `json:"vcard"`
}

// Dice represents an animated emoji that displays a random value.
type Dice struct {
	// Emoji on which the dice throw animation is based
	Emoji string `json:"phone_number"`
	// Value of the dice, 1-6 for “🎲” and “🎯” base emoji, 1-5
	// for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji
	Value string `json:"first_name"`
}

// PollOption contains information about one answer option in a poll.
type PollOption struct {
	// Option text, 1-100 characters
	Text string `json:"text"`
	// Number of users that voted for this option
	VoterCount string `json:"voter_count"`
}

// PollAnswer represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	// PollID is a unique poll identifier
	PollID string `json:"poll_id"`
	// The user, who changed the answer to the poll
	User *User `json:"user"`
	// 0-based identifiers of answer options, chosen by the user.
	// May be empty if the user retracted their vote.
	OptionIDs []int `json:"option_ids"`
}

// Poll contains information about a poll.
type Poll struct {
	// ID is a unique poll identifier
	ID string `json:"id"`
	// Question is a poll question, 1-255 characters
	Question string `json:"question"`
	// Options is list of poll options
	Options *[]PollOption `json:"options"`
	// TotalVoterCount is a total number of users that voted in the poll
	TotalVoterCount int `json:"total_voter_count"`
	// IsClosed is true, if the poll is closed
	IsClosed bool `json:"is_closed"`
	// IsAnonymous is true, if the poll is anonymous
	IsAnonymous bool `json:"is_anonymous"`
	// Type is poll type, currently can be “regular” or “quiz”
	Type string `json:"type"`
	// AllowsMultipleAnswers is true, if the poll allows multiple answers
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`
	// CorrectOptionID 0-based identifier of the correct answer option.
	// Available only for polls in the quiz mode, which are closed,
	// or was sent (not forwarded) by the bot or to the private chat with the bot.
	// Optional
	CorrectOptionID int `json:"correct_option_id"`
	// Explanation is a text that is shown when a user chooses an incorrect answer
	// or taps on the lamp icon in a quiz-style poll, 0-200 characters
	// Optional
	Explanation string `json:"explanation"`
	// ExplanationEntities are special entities
	// like usernames, URLs, bot commands, etc. that appear in the explanation
	// Optional
	ExplanationEntities *[]MessageEntity `json:"explanation_entities"`
	// OpenPeriod is amount of time in seconds the poll will be active after creation
	// Optional
	OpenPeriod int `json:"open_period"`
	// CloseDate isPoint in time (Unix timestamp) when the poll will be automatically closed
	// Optional
	CloseDate int `json:"close_date"`
}

// Location represents a point on the map.
type Location struct {
	// Longitude as defined by sender
	Longitude float64 `json:"longitude"`
	// Latitude as defined by sender
	Latitude float64 `json:"latitude"`
	// HorizontalAccuracy is the radius of uncertainty for the location, measured in meters; 0-1500
	// Optional
	HorizontalAccuracy float64 `json:"horizontal_accuracy"`
	// LivePeriod is a time relative to the message sending date, during which
	// the location can be updated, in seconds.
	// For active live locations only.
	// Optional
	LivePeriod int `json:"live_period"`
	// Heading is a direction in which user is moving, in degrees; 1-360.
	// For active live locations only.
	// Optional
	Heading int `json:"heading"`
	// ProximityAlertRadius is a maximum distance for proximity alerts about
	// approaching another chat member, in meters. For sent live locations only.
	// Optional
	ProximityAlertRadius int `json:"proximity_alert_radius"`
}

// Venue contains information about a venue, including its Location.
type Venue struct {
	// Location venue location. Can't be a live location
	Location *Location `json:"location"`
	// Title name of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// FoursquareID foursquare identifier of the venue
	// Optional
	FoursquareID string `json:"foursquare_id"`
	// FoursquareType Foursquare type of the venue.
	// (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	// Optional
	FoursquareType string `json:"foursquare_type"`
	// GooglePlaceID Google Places identifier of the venue
	// Optional
	GooglePlaceID string `json:"google_place_id"`
	// GooglePlaceType Google Places type of the venue
	// Optional
	GooglePlaceType string `json:"google_place_type"`
}

// ProximityAlertTriggered represents the content of a service message,
// sent whenever a user in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	// Traveler is user that triggered the alert
	Traveler *User `json:"traveler"`
	// Watcher is user that set the alert
	Watcher *User `json:"watcher"`
	// Distance is a distance between the users
	Distance int `json:"distance"`
}

// UserProfilePhotos contains a set of user profile photos.
type UserProfilePhotos struct {
	// TotalCount total number of profile pictures the target user has
	TotalCount int `json:"total_count"`
	// Photos requested profile pictures (in up to 4 sizes each)
	Photos *[][]PhotoSize `json:"photos"`
}

// File contains information about a file to download from Telegram.
type File struct {
	// FileID identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// FileUniqueID unique identifier for this file,
	//which is supposed to be the same over time and for different bots.
	//Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// FileSize file size, if known
	// Optional
	FileSize int `json:"file_size"`
	// FilePath file path.
	// Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
	// Optional
	FilePath string `json:"file_path"`
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options
// (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	// Keyboard is an array of button rows, each represented by an Array of KeyboardButton objects
	Keyboard *[][]KeyboardButton `json:"keyboard"`
	// ResizeKeyboard requests clients to resize the keyboard vertically for optimal fit
	// (e.g., make the keyboard smaller if there are just two rows of buttons).
	// Defaults to false, in which case the custom keyboard
	// is always of the same height as the app's standard keyboard.
	// Optional
	ResizeKeyboard bool `json:"resize_keyboard"`
	// OneTimeKeyboard requests clients to hide the keyboard as soon as it's been used.
	// The keyboard will still be available, but clients will automatically display
	// the usual letter-keyboard in the chat – the user can press a special button
	// in the input field to see the custom keyboard again.
	// Defaults to false.
	// Optional
	OneTimeKeyboard bool `json:"one_time_keyboard"`
	// Selective use this parameter if you want to show the keyboard to specific users only.
	// Targets:
	//  1) users that are @mentioned in the text of the Message object;
	//  2) if the bot's message is a reply (has Message.ReplyToMessage not nil), sender of the original message.
	// Example: A user requests to change the bot's language,
	// bot replies to the request with a keyboard to select the new language.
	// Other users in the group don't see the keyboard.
	// Optional
	Selective bool `json:"selective"`
}

// KeyboardButton represents one button of the reply keyboard.
//For simple text buttons String can be used instead of this object to specify text of the button.
//Optional fields request_contact, request_location, and request_poll are mutually exclusive.
type KeyboardButton struct {
	// Text of the button. If none of the optional fields are used,
	// it will be sent as a message when the button is pressed.
	Text string `json:"text"`
	// RequestContact if True, the user's phone number will be sent
	// as a contact when the button is pressed.
	// Available in private chats only.
	// Optional
	RequestContact bool `json:"request_contact"`
	// RequestLocation if True, the user's current location will be sent when the button is pressed.
	// Available in private chats only.
	// Optional
	RequestLocation bool `json:"request_location"`
	//  If specified, the user will be asked to create a poll
	// and send it to the bot when the button is pressed.
	// Available in private chats only
	// Optional
	RequestPoll bool `json:"request_poll"`
}

// KeyboardButtonPollType represents type of a poll,
// which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	// If quiz is passed, the user will be allowed to create only polls in the quiz mode.
	// If regular is passed, only regular polls will be allowed.
	// Otherwise, the user will be allowed to create a poll of any type.
	// Optional
	Type string `json:"type"`
}

// ReplyKeyboardRemove allows the Bot to hide a custom keyboard.
type ReplyKeyboardRemove struct {
	// RemoveKeyboard requests clients to remove the custom keyboard
	// (user will not be able to summon this keyboard;
	// if you want to hide the keyboard from sight but keep it accessible,
	// use one_time_keyboard in ReplyKeyboardMarkup).
	RemoveKeyboard bool `json:"remove_keyboard"`
	// Selective use this parameter if you want to remove the keyboard for specific users only.
	// Targets:
	//  1) users that are @mentioned in the text of the Message object;
	//  2) if the bot's message is a reply (has Message.ReplyToMessage not nil), sender of the original message.
	// Example: A user votes in a poll, bot returns confirmation message
	// in reply to the vote and removes the keyboard for that user,
	// while still showing the keyboard with poll options to users who haven't voted yet.
	// Optional
	Selective bool `json:"selective"`
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	// InlineKeyboard array of button rows, each represented by an Array of InlineKeyboardButton objects
	InlineKeyboard *[][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton represents one button of an inline keyboard.
// You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	// Text label text on the button
	Text string `json:"text"`
	// URL HTTP or tg:// url to be opened when button is pressed.
	// Optional
	URL string `json:"url"`
	// An HTTP URL used to automatically authorize the user.
	// Can be used as a replacement for the Telegram Login Widget.
	// Optional
	LoginURL *LoginURL `json:"login_url"`
	// CallbackData data to be sent in a callback query to the bot when button is pressed, 1-64 bytes.
	// Optional
	CallbackData string `json:"callback_data"`
	// SwitchInlineQuery if set, pressing the button will prompt the user to select one of their chats,
	// open that chat and insert the bot's username and the specified inline query in the input field.
	// Can be empty, in which case just the bot's username will be inserted.
	// This offers an easy way for users to start using your bot
	// in inline mode when they are currently in a private chat with it.
	// Especially useful when combined with switch_pm… actions – in this case
	// the user will be automatically returned to the chat they switched from,
	// skipping the chat selection screen.
	// Optional
	SwitchInlineQuery string `json:"switch_inline_query"`
	// SwitchInlineQueryCurrentChat if set, pressing the button will insert the bot's username
	// and the specified inline query in the current chat's input field.
	// Can be empty, in which case only the bot's username will be inserted.
	// This offers a quick way for the user to open your bot in inline mode
	// in the same chat – good for selecting something from multiple options.
	// Optional
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat"`
	// CallbackGame description of the game that will be launched when the user presses the button.
	// Optional
	CallbackGame *CallbackGame `json:"callback_game"`
	// Pay specify True, to send a Pay button.
	// NOTE: This type of button must always be the first button in the first row.
	// Optional
	Pay bool `json:"pay"`
}

// LoginURL represents a parameter of the inline keyboard button used to automatically authorize a user.
// Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram.
// All the user needs to do is tap/click a button and confirm that they want to log in:
type LoginURL struct {
	// An HTTP URL to be opened with user authorization data added to the query string when the button is pressed.
	// If the user refuses to provide authorization data, the original URL without information about the user will be opened.
	// The data added is the same as described in Receiving authorization data.
	URL string `json:"url"`
	// New text of the button in forwarded messages.
	// Optional
	ForwardText string `json:"forward_text"`
	// Username of a bot, which will be used for user authorization.
	// See Setting up a bot for more details. If not specified, the current bot's username will be assumed.
	// The url's domain must be the same as the domain linked with the bot.
	// See Linking your domain to the bot for more details.
	// Optional
	BotUsername string `json:"bot_username"`
	// Pass True to request the permission for your bot to send messages to the user.
	// Optional
	RequestWriteAccess bool `json:"request_write_access"`
}

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard.
type CallbackQuery struct {
	// ID unique identifier for this query
	ID string `json:"id"`
	// From sender
	From *User `json:"from"`
	// Message with the callback button that originated the query.
	// Note that message content and message date will not be available if the message is too old.
	// Optional
	Message *Message `json:"message"`
	// InlineMessageID identifier of the message sent via the bot in inline mode, that originated the query.
	// Optional
	InlineMessageID string `json:"inline_message_id"`
	// ChatInstance global identifier, uniquely corresponding to the chat to which
	// the message with the callback button was sent. Useful for high scores in games.
	ChatInstance string `json:"chat_instance"`
	// Data associated with the callback button. Be aware that
	// a bad client can send arbitrary data in this field.
	// Optional
	Data string `json:"data"`
	// GameShortName short name of a Game to be returned, serves as the unique identifier for the game.
	// Optional
	GameShortName string `json:"game_short_name"`
}

// ForceReply allows the Bot to have users directly reply to it without
// additional interaction.
type ForceReply struct {
	// ForceReply shows reply interface to the user,
	// as if they manually selected the bot's message and tapped 'Reply'.
	ForceReply bool `json:"force_reply"`
	// Selective use this parameter if you want to force reply from specific users only.
	// Targets:
	//  1) users that are @mentioned in the text of the Message object;
	//  2) if the bot's message is a reply (has Message.ReplyToMessage not nil), sender of the original message.
	// Optional
	Selective bool `json:"selective"`
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	// SmallFileID is a file identifier of small (160x160) chat photo.
	// This file_id can be used only for photo download and
	// only for as long as the photo is not changed.
	SmallFileID string `json:"small_file_id"`
	// Unique file identifier of small (160x160) chat photo,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	SmallFileUniqueID string `json:"small_file_unique_id"`
	// BigFileID is a file identifier of big (640x640) chat photo.
	// This file_id can be used only for photo download and
	// only for as long as the photo is not changed.
	BigFileID string `json:"big_file_id"`
	// Unique file identifier of big (640x640) chat photo,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	BigFileUniqueID string `json:"big_file_unique_id"`
}

// ChatMember contains information about one member of a chat.
type ChatMember struct {
	// User information about the user
	User *User `json:"user"`
	// Status the member's status in the chat.
	// Can be
	//  “creator”,
	//  “administrator”,
	//  “member”,
	//  “restricted”,
	//  “left” or
	//  “kicked”
	Status string `json:"status"`
	// CustomTitle owner and administrators only. Custom title for this user
	// Optional
	CustomTitle string `json:"custom_title"`
	// Owner and administrators only.
	// True, if the user's presence in the chat is hidden
	// Optional
	IsAnonymous bool `json:"is_anonymous"`
	// CanBeEdited administrators only.
	// True, if the bot is allowed to edit administrator privileges of that user.
	// Optional
	CanBeEdited bool `json:"can_be_edited"`
	// CanPostMessages administrators only.
	// True, if the administrator can post in the channel;
	// channels only
	// Optional
	CanPostMessages bool `json:"can_post_messages"`
	// CanEditMessages administrators only.
	// True, if the administrator can edit messages of other users and can pin messages;
	// channels only.
	// Optional
	CanEditMessages bool `json:"can_edit_messages"`
	// CanDeleteMessages administrators only.
	// True, if the administrator can delete messages of other users.
	// Optional
	CanDeleteMessages bool `json:"can_delete_messages"`
	// CanRestrictMembers administrators only.
	// True, if the administrator can restrict, ban or unban chat members.
	// Optional
	CanRestrictMembers bool `json:"can_restrict_members"`
	// CanPromoteMembers administrators only.
	// True, if the administrator can add new administrators
	// with a subset of their own privileges or demote administrators that he has promoted,
	// directly or indirectly (promoted by administrators that were appointed by the user).
	// Optional
	CanPromoteMembers bool `json:"can_promote_members"`
	// CanChangeInfo administrators and restricted only.
	// True, if the user is allowed to change the chat title, photo and other settings.
	// Optional
	CanChangeInfo bool `json:"can_change_info"`
	// CanInviteUsers administrators and restricted only.
	// True, if the user is allowed to invite new users to the chat.
	// Optional
	CanInviteUsers bool `json:"can_invite_users"`
	// CanPinMessages administrators and restricted only.
	// True, if the user is allowed to pin messages;
	// groups and supergroups only
	// Optional
	CanPinMessages bool `json:"can_pin_messages"`
	// IsMember restricted only.
	// True, if the user is a member of the chat at the moment of the request
	// Optional
	IsMember bool `json:"is_member"`
	// CanSendMessages restricted only.
	// True, if the user is allowed to send text messages, contacts, locations and venues
	// Optional
	CanSendMessages bool `json:"can_send_messages"`
	// CanSendMediaMessages restricted only.
	// True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes
	// Optional
	CanSendMediaMessages bool `json:"can_send_media_messages"`
	// CanSendPolls restricted only.
	// True, if the user is allowed to send polls
	// Optional
	CanSendPolls bool `json:"can_send_polls"`
	// CanSendOtherMessages restricted only.
	// True, if the user is allowed to send audios, documents,
	// photos, videos, video notes and voice notes.
	// Optional
	CanSendOtherMessages bool `json:"can_send_other_messages"`
	// CanAddWebPagePreviews restricted only.
	// True, if the user is allowed to add web page previews to their messages.
	// Optional
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	// UntilDate restricted and kicked only.
	// Date when restrictions will be lifted for this user;
	// unix time.
	// Optional
	UntilDate int `json:"until_date"`
}

// ChatPermissions describes actions that
// a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	// CanSendMessages is true, if the user is allowed to send
	// text messages, contacts, locations and venues
	// Optional
	CanSendMessages bool `json:"can_send_messages"`
	// CanSendMediaMessages is true, if the user is allowed to send audios,
	// documents, photos, videos, video notes and voice notes, implies can_send_messages
	// Optional
	CanSendMediaMessages bool `json:"can_send_media_messages"`
	// CanSendPolls is true, if the user is allowed to send polls,
	// implies can_send_messages
	// Optional
	CanSendPolls bool `json:"can_send_polls"`
	// CanSendOtherMessages is true, if the user is allowed to send animations,
	// games, stickers and use inline bots, implies can_send_media_messages
	// Optional
	CanSendOtherMessages bool `json:"can_send_other_messages"`
	// CanAddWebPagePreviews is True, if the user is allowed
	// to add web page previews to their messages,
	// implies can_send_media_messages
	// Optional
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	// CanChangeInfo is true, if the user is allowed to change the chat title,
	// photo and other settings. Ignored in public supergroups
	// Optional
	CanChangeInfo bool `json:"can_change_info"`
	// CanInviteUsers is true, if the user is allowed to invite new users to the chat
	// Optional
	CanInviteUsers bool `json:"can_invite_users"`
	// CanPinMessages is true, if the user is allowed to pin messages.
	// Ignored in public supergroups
	// Optional
	CanPinMessages bool `json:"can_pin_messages"`
}

// ChatLocation represents a location to which a chat is connected.
type ChatLocation struct {
	// Location to which the supergroup is connected.
	// Can't be a live location.
	Location *Location `json:"location"`
	// Location address; 1-64 characters, as defined by the chat owner
	Address string `json:"address"`
}

// BotCommand represents a bot command.
type BotCommand struct {
	// Text of the command, 1-32 characters.
	// Can contain only lowercase English letters, digits and underscores.
	Command string `json:"command"`
	// Description of the command, 3-256 characters.
	Description string `json:"description"`
}

// ResponseParameters contains information about why a request was unsuccessful.
type ResponseParameters struct {
	// The group has been migrated to a supergroup with the specified identifier.
	// This number may be greater than 32 bits and some programming languages may have
	// difficulty/silent defects in interpreting it.
	// But it is smaller than 52 bits, so a signed 64 bit integer or
	// double-precision float type are safe for storing this identifier.
	// Optional
	MigrateToChatID int `json:"migrate_to_chat_id"`
	// In case of exceeding flood control, the number of seconds left to wait
	// before the request can be repeated
	// Optional
	RetryAfter int `json:"retry_after"`
}

// InputFile   represents the contents of a file to be uploaded. 
// Must be posted using multipart/form-data in the usual way 
// that files are uploaded via the browser.
type InputFile interface {}

// InputMediaPhoto represents a photo to be sent.
type InputMediaPhoto struct {
	// Type of the result, must be photo.
	Type string `json:"type"`
	// Media file to send. Pass a file_id to send a file that
	// exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass “attach://<file_attach_name>” to upload a new one
	// using multipart/form-data under <file_attach_name> name.
	Media string `json:"media"`
	// Caption of the photo to be sent, 0-1024 characters after entities parsing.
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the photo caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// CaptionEntities is a List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
}

// InputMediaVideo represents a video to be sent.
type InputMediaVideo struct {
	// Type of the result, must be video.
	Type string `json:"type"`
	// Media file to send. Pass a file_id to send a file
	// that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass “attach://<file_attach_name>” to upload a new one
	// using multipart/form-data under <file_attach_name> name.
	Media string `json:"media"`
	// Thumbnail of the file sent;
	// can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file,
	// so you can pass “attach://<file_attach_name>” if the
	// thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	// Optional
	Thumb *InputFile `json:"thumb"`
	// Caption of the video to be sent, 0-1024 characters after entities parsing.
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the video caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// CaptionEntities is a list of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// Width video width
	// Optional
	Width int `json:"width"`
	// Height video height
	// Optional
	Height int `json:"height"`
	// Duration video duration
	// Optional
	Duration int `json:"duration"`
	// SupportsStreaming pass True, if the uploaded video is suitable for streaming.
	// Optional
	SupportsStreaming bool `json:"supports_streaming"`
}

// InputMediaAnimation represents an animation file
// (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	// Type of the result, must be animation.
	Type string `json:"type"`
	// Media file to send. Pass a file_id to send a file
	// that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass “attach://<file_attach_name>” to upload a new one
	// using multipart/form-data under <file_attach_name> name.
	Media string `json:"media"`
	// Thumbnail of the file sent;
	// can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file,
	// so you can pass “attach://<file_attach_name>” if the
	// thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	// Optional
	Thumb *InputFile `json:"thumb"`
	// Caption of the animation to be sent, 0-1024 characters after entities parsing.
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the animation caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// CaptionEntities is a list of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// Width animation width
	// Optional
	Width int `json:"width"`
	// Height animation height
	// Optional
	Height int `json:"height"`
	// Duration animation duration
	// Optional
	Duration int `json:"duration"`
}

// InputMediaAudio represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	// Type of the result, must be audio.
	Type string `json:"type"`
	// Media file to send. Pass a file_id to send a file
	// that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass “attach://<file_attach_name>” to upload a new one
	// using multipart/form-data under <file_attach_name> name.
	Media string `json:"media"`
	// Thumbnail of the file sent;
	// can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file,
	// so you can pass “attach://<file_attach_name>” if the
	// thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	// Optional
	Thumb *InputFile `json:"thumb"`
	// Caption of the audio to be sent, 0-1024 characters after entities parsing.
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the animation caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// CaptionEntities is a list of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// Duration of the audio in seconds
	// Optional
	Duration int `json:"duration"`
	// Performer of the audio
	// Optional
	Performer string `json:"performer"`
	// Title of the audio
	// Optional
	Title string `json:"title"`
}

// InputMediaDocument represents a general file to be sent.
type InputMediaDocument struct {
	// Type of the result, must be document.
	Type string `json:"type"`
	// Media file to send. Pass a file_id to send a file
	// that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass “attach://<file_attach_name>” to upload a new one
	// using multipart/form-data under <file_attach_name> name.
	Media string `json:"media"`
	// Thumbnail of the file sent;
	// can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file,
	// so you can pass “attach://<file_attach_name>” if the
	// thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	// Optional
	Thumb *InputFile `json:"thumb"`
	// Caption of the document to be sent, 0-1024 characters after entities parsing.
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the animation caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// CaptionEntities is a list of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// Disables automatic server-side content type detection for
	// files uploaded using multipart/form-data.
	// Always true, if the document is sent as part of an album.
	// Optional
	DisableContentTypeDetection bool `json:"disable_content_type_detection"`
}