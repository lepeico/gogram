package telegram

import (
	"encoding/json"
	"errors"

	"github.com/lepeico/gogram/internal/client"
)

type Telegram struct {
	client.Client
}

func New(token string) *Telegram {
	return &Telegram{*client.NewClient(token)}
}

func (tg *Telegram) GetMe() (User, error) {
	var bot User

	response, err := tg.Call("getMe", nil)
	if err != nil {
		return bot, err
	}
	json.Unmarshal(response, &bot)

	return bot, nil
}

func (tg *Telegram) SendMessage(chatID, text string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("text", text)

	res, err = tg.Call("sendMessage", message)
	return
}

func (tg *Telegram) SendDocument(chatID string, document interface{}) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("document", document)

	res, err = tg.Call("sendDocument", message)
	return
}

func (tg *Telegram) ForwardMessage(chatID, fromChatID, messageID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("from_chat_id", fromChatID)
	message.Set("message_id", messageID)

	res, err = tg.Call("forwardMessage", message)
	return
}

func (tg *Telegram) CopyMessage(chatID, fromChatID, messageID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("from_chat_id", fromChatID)
	message.Set("message_id", messageID)

	res, err = tg.Call("copyMessage", message)
	return
}

func (tg *Telegram) SendPhoto(chatID string, photo interface{}) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("photo", photo)

	res, err = tg.Call("sendPhoto", message)
	return
}

func (tg *Telegram) SendAudio(chatID string, audio interface{}) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("audio", audio)

	res, err = tg.Call("sendAudio", message)
	return
}

func (tg *Telegram) SendVideo(chatID string, video interface{}) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("video", video)

	res, err = tg.Call("sendVideo", message)
	return
}

func (tg *Telegram) SendAnimation(chatID string, animation interface{}) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("animation", animation)

	res, err = tg.Call("sendAnimation", message)
	return
}

func (tg *Telegram) SendVoice(chatID string, voice interface{}) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("voice", voice)

	res, err = tg.Call("sendVoice", message)
	return
}

func (tg *Telegram) SendVideoNote(chatID string, videoNote interface{}) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("video_note", videoNote)

	res, err = tg.Call("sendVideoNote", message)
	return
}

//mediaGroup - array of audios, documents, photos, videos and links to the media
// func (tg *Telegram) SendMediaGroup(chatID string, mediaGroup []interface{}) (res json.RawMessage, err error) {
// 	if len(mediaGroup) < 2 || len(mediaGroup) > 10 {
// 		err = errors.New("mediaGroup size is out of range")
// 		return nil, err
// 	}
//
// 	message := client.Payload{}
// 	message.Set("chat_id", chatID)
// 	message.Set("media", mediaGroup)
// 	//log.Println(media)
//
// 	res, err = tg.Call("sendMediaGroup", message)
// 	return
// }

func (tg *Telegram) SendLocation(chatID string, latitude, longitude float32) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("latitude", latitude)
	message.Set("longitude", longitude)

	res, err = tg.Call("sendLocation", message)
	return
}

func (tg *Telegram) EditMessageLiveLocation(chatID, messageID string, latitude, longitude float32) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("message_id", messageID)
	message.Set("latitude", latitude)
	message.Set("longitude", longitude)

	res, err = tg.Call("editMessageLiveLocation", message)
	return
}

func (tg *Telegram) StopMessageLiveLocation(chatID, messageID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("message_id", messageID)

	res, err = tg.Call("stopMessageLiveLocation", message)
	return
}

func (tg *Telegram) SendVenue(chatID string, latitude, longitude float32, title, address string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("latitude", latitude)
	message.Set("longitude", longitude)
	message.Set("title", title)
	message.Set("address", address)

	res, err = tg.Call("sendVenue", message)
	return
}

func (tg *Telegram) SendContact(chatID string, phoneNumber, firstName string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("phone_number", phoneNumber)
	message.Set("first_name", firstName)

	res, err = tg.Call("sendContact", message)
	return
}

func (tg *Telegram) SendPoll(chatID string, question string, options []string) (res json.RawMessage, err error) {
	if len(options) < 2 || len(options) > 10 {
		err = errors.New("mediaGroup size is out of range")
		return nil, err
	}

	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("question", question)
	message.Set("options", options)

	res, err = tg.Call("sendPoll", message)
	return
}

func (tg *Telegram) SendDice(chatID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)

	res, err = tg.Call("sendDice", message)
	return
}

func (tg *Telegram) SendChatAction(chatID, action string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("action", action)

	res, err = tg.Call("sendChatAction", message)
	return
}

func (tg *Telegram) GetUserProfilePhotos(userID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("user_id", userID)

	res, err = tg.Call("getUserProfilePhotos", message)
	return
}

func (tg *Telegram) GetFile(fileID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("file_id", fileID)

	res, err = tg.Call("getFile", message)
	return
}

//TODO: test
func (tg *Telegram) KickChatMember(chatID, userID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("user_id", userID)

	res, err = tg.Call("kickChatMember", message)
	return
}

//TODO: test
func (tg *Telegram) UnbanChatMember(chatID, userID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("user_id", userID)

	res, err = tg.Call("unbanChatMember", message)
	return
}

//TODO: test
func (tg *Telegram) RestrictChatMember(chatID, userID string, permissions ChatPermissions) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("user_id", userID)
	message.Set("permissions", permissions)

	res, err = tg.Call("restrictChatMember", message)
	return
}

//TODO: test, add optional parameters
func (tg *Telegram) PromoteChatMember(chatID, userID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("user_id", userID)

	res, err = tg.Call("promoteChatMember", message)
	return
}

//TODO: test
func (tg *Telegram) SetChatAdministratorCustomTitle(chatID, userID, customTitle string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("user_id", userID)
	message.Set("custom_title", customTitle)

	res, err = tg.Call("setChatAdministratorCustomTitle", message)
	return
}

//TODO: test
func (tg *Telegram) SetChatPermissions(chatID string, permissions ChatPermissions) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("permissions", permissions)

	res, err = tg.Call("setChatPermissions", message)
	return
}

//TODO: test
func (tg *Telegram) ExportChatInviteLink(chatID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)

	res, err = tg.Call("exportChatInviteLink", message)
	return
}

//TODO: test
func (tg *Telegram) SetChatPhoto(chatID string, photo interface{}) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("photo", photo)

	res, err = tg.Call("setChatPhoto", message)
	return
}

//TODO: test
func (tg *Telegram) DeleteChatPhoto(chatID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)

	res, err = tg.Call("deleteChatPhoto", message)
	return
}

//TODO: test
func (tg *Telegram) SetChatTitle(chatID, title string) (res json.RawMessage, err error) {
	if len(title) > 255 {
		err = errors.New("title size is out of range")
		return nil, err
	}
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("title", title)

	res, err = tg.Call("setChatTitle", message)
	return
}

//TODO: test, description is optional
func (tg *Telegram) SetChatDescription(chatID, description string) (res json.RawMessage, err error) {
	if len(description) > 255 {
		err = errors.New("description size is out of range")
		return nil, err
	}
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("description", description)

	res, err = tg.Call("setChatDescription", message)
	return
}

func (tg *Telegram) PinChatMessage(chatID, messageID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("message_id", messageID)

	res, err = tg.Call("pinChatMessage", message)
	return
}

func (tg *Telegram) UnpinChatMessage(chatID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)

	res, err = tg.Call("unpinChatMessage", message)
	return
}

func (tg *Telegram) UnpinAllChatMessages(chatID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)

	res, err = tg.Call("unpinAllChatMessages", message)
	return
}

//TODO: test
func (tg *Telegram) LeaveChat(chatID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)

	res, err = tg.Call("leaveChat", message)
	return
}

//TODO: test
func (tg *Telegram) GetChatAdministrators(chatID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)

	res, err = tg.Call("getChatAdministrators", message)
	return
}

func (tg *Telegram) GetChatMembersCount(chatID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)

	res, err = tg.Call("getChatMembersCount", message)
	return
}

func (tg *Telegram) GetChatMember(chatID, userID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("user_id", userID)

	res, err = tg.Call("getChatMember", message)
	return
}

//TODO: test
func (tg *Telegram) SetChatStickerSet(chatID, stickerSetName string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("sticker_set_name", stickerSetName)

	res, err = tg.Call("setChatStickerSet", message)
	return
}

//TODO: test
func (tg *Telegram) DeleteChatStickerSet(chatID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)

	res, err = tg.Call("deleteChatStickerSet", message)
	return
}

//TODO: test
func (tg *Telegram) AnswerCallbackQuery(callbackQueryID string) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("callback_query_id", callbackQueryID)

	res, err = tg.Call("answerCallbackQuery", message)
	return
}

//TODO: test
func (tg *Telegram) SetMyCommands(commands []BotCommand) (res json.RawMessage, err error) {
	message := client.Payload{}
	message.Set("commands", commands)

	res, err = tg.Call("setMyCommands", message)
	return
}

func (tg *Telegram) GetMyCommands() (res json.RawMessage, err error) {
	message := client.Payload{}

	res, err = tg.Call("getMyCommands", message)
	return
}
