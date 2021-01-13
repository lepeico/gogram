package telegram

// InlineQuery represents an incoming inline query.
// When the user sends an empty query,
// your bot could return some default or trending results.
type InlineQuery struct {
	// Unique identifier for this query
	ID string `json:"id"`
	// From is a sender
	From *User `json:"from"`
	// Sender location, only for bots that request user location
	// Optional
	Location *Location `json:"location"`
	// Text of the query (up to 256 characters)
	Query string `json:"query"`
	// Offset of the results to be returned, can be controlled by the bot
	Offset string `json:"offset"`
}

// InlineQueryResultArticle is an inline query response article.
type InlineQueryResultArticle struct {
	// Type of the result, must be article.
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 Bytes.
	ID string `json:"id"`
	// Title of the result
	Title string `json:"title"`
	// InputMessageContent content of the message to be sent.
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// ReplyMarkup Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// URL of the result.
	// Optional
	URL string `json:"url"`
	// HideURL pass True, if you don't want the URL to be shown in the message.
	// Optional
	HideURL bool `json:"hide_url"`
	// Description short description of the result.
	// Optional
	Description string `json:"description"`
	// ThumbURL url of the thumbnail for the result
	// Optional
	ThumbURL string `json:"thumb_url"`
	// ThumbWidth thumbnail width
	// Optional
	ThumbWidth int `json:"thumb_width"`
	// ThumbHeight thumbnail height
	// Optional
	ThumbHeight int `json:"thumb_height"`
}

// InlineQueryResultPhoto represents a link to a photo.
// By default, this photo will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	// Type of the result, must be article.
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 Bytes.
	ID string `json:"id"`
	// URL a valid URL of the photo. Photo must be in jpeg format.
	// Photo size must not exceed 5MB.
	PhotoURL string `json:"photo_url"`
	// ThumbURL url of the thumbnail for the photo.
	// Optional
	ThumbURL string `json:"thumb_url"`
	// Width of the photo
	// Optional
	PhotoWidth int `json:"photo_width"`
	// Height of the photo
	// Optional
	PhotoHeight int `json:"photo_height"`
	// Title for the result
	// Optional
	Title string `json:"title"`
	// Description short description of the result
	// Optional
	Description string `json:"description"`
	// Caption of the photo to be sent, 0-1024 characters after entities parsing.
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the photo caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// ReplyMarkup inline keyboard attached to the message.
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the photo.
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// InlineQueryResultGIF is an inline query response GIF.
type InlineQueryResultGIF struct {
	// Type of the result, must be gif.
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`
	// URL a valid URL for the GIF file. File size must not exceed 1MB.
	GifURL string `json:"gif_url"`
	// Width of the GIF
	// Optional
	GifWidth int `json:"gif_width"`
	// Height of the GIF
	// Optional
	GifHeight int `json:"gif_height"`
	// Duration of the GIF
	// Optional
	GifDuration int `json:"gif_duration"`
	// ThumbURL url of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result.
	ThumbURL string `json:"thumb_url"`
	// MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”,
	// or “video/mp4”. Defaults to “image/jpeg”
	// Optional
	ThumbMimeType string `json:"thumb_mime_type"`
	// Title for the result
	// Optional
	Title string `json:"title"`
	// Caption of the GIF file to be sent, 0-1024 characters after entities parsing.
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// ReplyMarkup inline keyboard attached to the message
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the GIF animation.
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// InlineQueryResultMPEG4GIF represents a link to a video animation
// (H.264/MPEG-4 AVC video without sound). By default,
//this animated MPEG-4 file will be sent by the user with optional caption.
type InlineQueryResultMPEG4GIF struct {
	// Type of the result, must be mpeg4_gif
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes
	ID string `json:"id"`
	// URL a valid URL for the MP4 file. File size must not exceed 1MB
	Mpeg4URL string `json:"mpeg4_url"`
	// Width video width
	// Optional
	Mpeg4Width int `json:"mpeg4_width"`
	// Height vVideo height
	// Optional
	Mpeg4Height int `json:"mpeg4_height"`
	// Duration video duration
	// Optional
	Mpeg4Duration int `json:"mpeg4_duration"`
	// ThumbURL url of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result.
	ThumbURL string `json:"thumb_url"`
	// MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”,
	// or “video/mp4”. Defaults to “image/jpeg”
	// Optional
	ThumbMimeType string `json:"thumb_mime_type"`
	// Title for the result
	// Optional
	Title string `json:"title"`
	// Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing.
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// ReplyMarkup inline keyboard attached to the message
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the video animation
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// InlineQueryResultVideo represents a link to a page containing
// an embedded video player or a video file.
// By default, this video file will be sent by the user with an optional caption.
type InlineQueryResultVideo struct {
	// Type of the result, must be video
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes
	ID string `json:"id"`
	// URL a valid url for the embedded video player or video file
	VideoURL string `json:"video_url"`
	// MimeType of the content of video url, “text/html” or “video/mp4”
	MimeType string `json:"mime_type"`
	// ThumbURL url of the thumbnail (jpeg only) for the video
	// optional
	ThumbURL string `json:"thumb_url"`
	// Title for the result
	Title string `json:"title"`
	// Caption of the video to be sent, 0-1024 characters after entities parsing
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the video caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// Width video width
	// Optional
	VideoWidth int `json:"video_width"`
	// Height video height
	// Optional
	VideoHeight int `json:"video_height"`
	// Duration video duration in seconds
	// Optional
	VideoDuration int `json:"video_duration"`
	// Description short description of the result
	// Optional
	Description string `json:"description"`
	// ReplyMarkup inline keyboard attached to the message
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the video.
	// This field is required if InlineQueryResultVideo is used to send
	// an HTML-page as a result (e.g., a YouTube video).
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// InlineQueryResultAudio represents a link to an MP3 audio file.
// By default, this audio file will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
	// Type of the result, must be audio
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes
	ID string `json:"id"`
	// URL a valid url for the audio file
	AudioURL string `json:"audio_url"`
	// Title is a title
	Title string `json:"title"`
	// Caption 0-1024 characters after entities parsing
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the audio caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// Performer is a performer
	// Optional
	Performer string `json:"performer"`
	// Duration audio duration in seconds
	// Optional
	AudioDuration int `json:"audio_duration"`
	// ReplyMarkup inline keyboard attached to the message
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the audio
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// InlineQueryResultVoice represents a link to a voice recording in an .OGG container encoded with OPUS.
// By default, this voice recording will be sent by the user.
// Alternatively, you can use input_message_content to send a message
// with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	// Type of the result, must be voice
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes
	ID string `json:"id"`
	// URL a valid URL for the voice recording
	VoiceURL string `json:"voice_url"`
	// Title recording title
	Title string `json:"title"`
	// Caption 0-1024 characters after entities parsing
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the voice caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// Duration recording duration in seconds
	// Optional
	VoiceDuration int `json:"voice_duration"`
	// ReplyMarkup inline keyboard attached to the message
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the voice recording
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// InlineQueryResultDocument represents a link to a file.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send
// a message with the specified content instead of the file.
// Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	// Type of the result, must be document
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes
	ID string `json:"id"`
	// Title for the result
	Title string `json:"title"`
	// Caption of the document to be sent, 0-1024 characters after entities parsing
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the voice caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// URL a valid url for the file
	DocumentURL string `json:"document_url"`
	// MimeType of the content of the file, either “application/pdf” or “application/zip”
	MimeType string `json:"mime_type"`
	// Description short description of the result
	// Optional
	Description string `json:"description"`
	// ReplyMarkup nline keyboard attached to the message
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the file
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// ThumbURL url of the thumbnail (jpeg only) for the file
	// Optional
	ThumbURL string `json:"thumb_url"`
	// ThumbWidth thumbnail width
	// Optional
	ThumbWidth int `json:"thumb_width"`
	// ThumbHeight thumbnail height
	// Optional
	ThumbHeight int `json:"thumb_height"`
}

// InlineQueryResultLocation represents a location on a map.
// By default, the location will be sent by the user.
// Alternatively, you can use input_message_content to send
// a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
	// Type of the result, must be location
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 Bytes
	ID string `json:"id"`
	// Latitude  of the location in degrees
	Latitude float64 `json:"latitude"`
	// Longitude of the location in degrees
	Longitude float64 `json:"longitude"`
	// Title of the location
	Title string `json:"title"`
	//  The radius of uncertainty for the location, measured in meters; 0-1500
	// Optional
	HorizontalAccuracy float64 `json:"horizontal_accuracy"`
	// Period in seconds for which the location can be updated, should be between 60 and 86400.
	// Optional
	LivePeriod int `json:"live_period"`
	// For live locations, a direction in which the user is moving, in degrees.
	// Must be between 1 and 360 if specified.
	// Optional
	Heading int `json:"heading"`
	// For live locations, a maximum distance for proximity alerts about approaching another chat member,
	// in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius"`
	// ReplyMarkup inline keyboard attached to the message
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the location
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// ThumbURL url of the thumbnail for the result
	// Optional
	ThumbURL string `json:"thumb_url"`
	// ThumbWidth thumbnail width
	// Optional
	ThumbWidth int `json:"thumb_width"`
	// ThumbHeight thumbnail height
	// Optional
	ThumbHeight int `json:"thumb_height"`
}

// InlineQueryResultVenue represents a venue.
// By default, the venue will be sent by the user.
// Alternatively, you can use input_message_content to send
// a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	// Type of the result, must be venue
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 Bytes]
	ID string `json:"id"`
	// Latitude of the venue location in degrees
	Latitude float64 `json:"latitude"`
	// Longitude of the venue location in degrees
	Longitude float64 `json:"longitude"`
	// Title of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// FoursquareID foursquare identifier of the venue if known
	// Optional
	FoursquareID string `json:"foursquare_id"`
	// FoursquareType foursquare type of the venue, if known.
	// (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	// Optional
	FoursquareType string `json:"foursquare_type"`
	// GooglePlaceID Google Places identifier of the venue
	// Optional
	GooglePlaceID string `json:"google_place_id"`
	// GooglePlaceTypeGoogle Places type of the venue
	// Optional
	GooglePlaceType string `json:"google_place_type"`
	// ReplyMarkup inline keyboard attached to the message
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the venue
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// ThumbURL url of the thumbnail for the result
	// Optional
	ThumbURL string `json:"thumb_url"`
	// ThumbWidth thumbnail width
	// Optional
	ThumbWidth int `json:"thumb_width"`
	// ThumbHeight thumbnail height
	// Optional
	ThumbHeight int `json:"thumb_height"`
}

// InlineQueryResultContact represents a contact with a phone number.
// By default, this contact will be sent by the user.
//  Alternatively, you can use input_message_content to send
//a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
	// Type of the result, must be contact
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 Bytes]
	ID string `json:"id"`
	// PhoneNumber contact's phone number
	PhoneNumber string `json:"phone_number"`
	// FirstName contact's first name
	FirstName string `json:"first_name"`
	// LastName contact's first name
	// Optional
	LastName string `json:"last_name"`
	// Vcard Additional data about the contact in the form of a vCard, 0-2048 bytes
	// Optional
	Vcard string `json:"vcard"`
	// ReplyMarkup inline keyboard attached to the message
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the venue
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// ThumbURL url of the thumbnail for the result
	// Optional
	ThumbURL string `json:"thumb_url"`
	// ThumbWidth thumbnail width
	// Optional
	ThumbWidth int `json:"thumb_width"`
	// ThumbHeight thumbnail height
	// Optional
	ThumbHeight int `json:"thumb_height"`
}

// InlineQueryResultGame is an inline query response game.
type InlineQueryResultGame struct {
	// Type of the result, must be game
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes
	ID string `json:"id"`
	// GameShortName short name of the game
	GameShortName string `json:"game_short_name"`
	// ReplyMarkup inline keyboard attached to the message
	// optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}

// InlineQueryResultCachedPhoto represents a link to a photo stored on the Telegram servers.
// By default, this photo will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send
// a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	// Type of the result, must be photo.
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`
	// PhotoID a valid file identifier of the photo.
	PhotoFileID string `json:"photo_file_id"`
	// Title for the result.
	// Optional
	Title string `json:"title"`
	// Description short description of the result.
	// Optional
	Description string `json:"description"`
	// Caption of the photo to be sent, 0-1024 characters after entities parsing.
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the voice caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// ReplyMarkup inline keyboard attached to the message.
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the photo.
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// InlineQueryResultCachedGIF represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send
// a message with specified content instead of the animation.
type InlineQueryResultCachedGIF struct {
	// Type of the result, must be gif.
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`
	// GifID a valid file identifier for the GIF file.
	GifFileID string `json:"gif_file_id"`
	// Title for the result
	Title string `json:"title"`
	// Caption of the GIF file to be sent, 0-1024 characters after entities parsing.
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// ReplyMarkup inline keyboard attached to the message.
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the GIF animation.
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// InlineQueryResultCachedMpeg4Gif is an inline query response with cached
// H.264/MPEG-4 AVC video without sound gif.
type InlineQueryResultCachedMpeg4Gif struct {
	// Type of the result, must be mpeg4_gif
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes
	ID string `json:"id"`
	// MGifID a valid file identifier for the MP4 file
	Mpeg4FilefID string `json:"mpeg4_file_id"`
	// Title for the result
	// Optional
	Title string `json:"title"`
	// Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing.
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// ReplyMarkup inline keyboard attached to the message.
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the video animation.
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// InlineQueryResultCachedSticker represents a link to a sticker stored on the Telegram servers.
// By default, this sticker will be sent by the user.
// Alternatively, you can use input_message_content to send
// a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	// Type of the result, must be sticker
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes
	ID string `json:"id"`
	// StickerID a valid file identifier of the sticker
	StickerFileID string `json:"sticker_file_id"`
	// ReplyMarkup inline keyboard attached to the message
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// InputMessageContent content of the message to be sent instead of the sticker
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedDocument is an inline query response with cached document.
type InlineQueryResultCachedDocument struct {
	// Type of the result, must be document
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes
	ID string `json:"id"`
	// DocumentID a valid file identifier for the file
	DocumentFileID string `json:"document_file_id"`
	// Title for the result
	Title string `json:"title"`
	// Description short description of the result
	// Optional
	Description string `json:"description"`
	// Caption of the document to be sent, 0-1024 characters after entities parsing
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the video caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// ReplyMarkup inline keyboard attached to the message
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// InputMessageContent content of the message to be sent instead of the file
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVideo represents a link to a video file stored on the Telegram servers.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send
// a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	// Type of the result, must be video
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes
	ID string `json:"id"`
	// VideoID a valid file identifier for the video file
	VideoFileID string `json:"video_file_id"`
	// Title for the result
	Title string `json:"title"`
	// Description short description of the result
	// Optional
	Description string `json:"description"`
	// Caption of the video to be sent, 0-1024 characters after entities parsing
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the video caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// ReplyMarkup inline keyboard attached to the message
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the video
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// InlineQueryResultCachedVoice represents a link to a voice message stored on the Telegram servers.
// By default, this voice message will be sent by the user.
// Alternatively, you can use input_message_content to send
// a message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
	// Type of the result, must be voice
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes
	ID string `json:"id"`
	// VoiceID a valid file identifier for the voice message
	VoiceFileID string `json:"voice_file_id"`
	// Title voice message title
	Title string `json:"title"`
	// Caption 0-1024 characters after entities parsing
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the video caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// ReplyMarkup inline keyboard attached to the message
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the voice message
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

// InlineQueryResultCachedAudio is an inline query response with cached audio.
type InlineQueryResultCachedAudio struct {
	// Type of the result, must be audio
	Type string `json:"type"`
	// ID unique identifier for this result, 1-64 bytes
	ID string `json:"id"`
	// AudioID a valid file identifier for the audio file
	AudioFileID string `json:"audio_file_id"`
	// Caption 0-1024 characters after entities parsing
	// Optional
	Caption string `json:"caption"`
	// ParseMode mode for parsing entities in the video caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// Optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	// Optional
	CaptionEntities *[]MessageEntity `json:"caption_entities"`
	// ReplyMarkup inline keyboard attached to the message
	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// InputMessageContent content of the message to be sent instead of the audio
	// Optional
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}
// InputMessageContent  represents the content of 
// a message to be sent as a result of an inline query. 
// Telegram clients currently support the following 4 types:
type InputMessageContent interface {}

// InputTextMessageContent represents the content of
// a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	// Text of the message to be sent, 1-4096 characters
	MessageText string `json:"message_text"`
	// ParseMode mode for parsing entities in the message text.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	// optional
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in message text,
	// which can be specified instead of parse_mode
	// Optional
	Entities *[]MessageEntity `json:"entities"`
	// DisableWebPagePreview disables link previews for links in the sent message
	// Optional
	DisableWebPagePreview bool `json:"disable_web_page_preview"`
}

// InputLocationMessageContent represents the content of
// a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	// Latitude of the location in degrees
	Latitude float64 `json:"latitude"`
	// Longitude of the location in degrees
	Longitude float64 `json:"longitude"`
	// HorizontalAccuracy the radius of uncertainty for the location,
	// measured in meters; 0-1500
	// Optional
	HorizontalAccuracy float64 `json:"horizontal_accuracy"`
	// LivePeriod period in seconds for which the location can be updated,
	// should be between 60 and 86400.
	// Optional
	LivePeriod int `json:"live_period"`
	// For live locations, a direction in which the user is moving, in degrees.
	// Must be between 1 and 360 if specified.
	// Optional
	Heading int `json:"heading"`
	// For live locations, a maximum distance for proximity alerts about approaching another chat member,
	// in meters. Must be between 1 and 100000 if specified.
	// Optional
	ProximityAlertRadius int `json:"proximity_alert_radius"`
}

// InputVenueMessageContent represents the content of a venue message
// to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	// Latitude of the venue in degrees
	Latitude float64 `json:"latitude"`
	// Longitude of the venue in degrees
	Longitude float64 `json:"longitude"`
	// Title name of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// FoursquareID foursquare identifier of the venue, if known
	// Optional
	FoursquareID string `json:"foursquare_id"`
	// FoursquareType foursquare type of the venue, if known.
	// (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	// Optional
	FoursquareType string `json:"foursquare_type"`
	// GooglePlaceID Google Places identifier of the venue
	// Optional
	GooglePlaceID string `json:"google_place_id"`
	// GooglePlaceTypeGoogle Places type of the venue
	// Optional
	GooglePlaceType string `json:"google_place_type"`
}

// InputContactMessageContent contains a contact for displaying
// as an inline query result.
type InputContactMessageContent struct {
	// 	PhoneNumber contact's phone number
	PhoneNumber string `json:"phone_number"`
	// FirstName contact's first name
	FirstName string `json:"first_name"`
	// LastName contact's last name
	// Optional
	LastName string `json:"last_name"`
	// Vcard additional data about the contact in the form of a vCard, 0-2048 bytes
	// Optional
	Vcard string `json:"vcard"`
}

// ChosenInlineResult represents a result of an inline query
// that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	// The unique identifier for the result that was chosen
	ResultID string `json:"result_id"`
	// The user that chose the result
	From *User `json:"from"`
	// Sender location, only for bots that require user location
	// Optional
	Location *Location `json:"location"`
	// Identifier of the sent inline message.
	// Available only if there is an inline keyboard attached to the message.
	// Will be also received in callback queries and can be used to edit the message.
	// Optional
	InlineMessageID string `json:"inline_message_id"`
	// The query that was used to obtain the result
	Query string `json:"query"`
}