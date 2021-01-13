package telegram

import (
	"github.com/lepeico/gogram/internal/client"
)

type fields map[string]interface{}

type Telegram struct {
	client client.Client
}

func New(token string) *Telegram {
	return &Telegram{*client.NewClient(token)}
}

func createPayload(f fields, opts []Option) (message client.Payload) {
	message = client.Payload(f)
	for _, opt := range opts {
		message.Set(opt.Name, opt.Parameter)
	}
	return
}

/// Available methods

func (tg *Telegram) GetMe() (bot User, err error) {
	err = tg.client.Call("getMe", nil, &bot)
	return
}

func (tg *Telegram) SendMessage(chat Chat, text string, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("sendMessage", createPayload(fields{
		"chat_id": chat.ID,
		"text":    text,
	}, opts), &msg)
	return
}

func (tg *Telegram) ForwardMessage(chat Chat, msg Message, opts ...Option) (forwardedMsg Message, err error) {
	err = tg.client.Call("forwardMessage", createPayload(fields{
		"chat_id":      chat.ID,
		"from_chat_id": msg.Chat.ID,
		"message_id":   msg.MessageID,
	}, opts), &forwardedMsg)
	return
}

func (tg *Telegram) CopyMessage(chat Chat, msg Message, opts ...Option) (msgID MessageID, err error) {
	err = tg.client.Call("copyMessage", createPayload(fields{
		"chat_id":      chat.ID,
		"from_chat_id": msg.Chat.ID,
		"message_id":   msg.MessageID,
	}, opts), msgID)
	return
}

func (tg *Telegram) SendPhoto(chat Chat, photo interface{}, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("sendPhoto", createPayload(fields{
		"chat_id": chat.ID,
		"photo":   photo,
	}, opts), &msg)
	return
}

func (tg *Telegram) SendAudio(chat Chat, audio interface{}, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("sendAudio", createPayload(fields{
		"chat_id": chat.ID,
		"audio":   audio,
	}, opts), &msg)
	return
}

func (tg *Telegram) SendDocument(chat Chat, document interface{}, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("sendDocument", createPayload(fields{
		"chat_id":  chat.ID,
		"document": document,
	}, opts), &msg)
	return
}

func (tg *Telegram) SendVideo(chat Chat, video interface{}, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("sendVideo", createPayload(fields{
		"chat_id": chat.ID,
		"video":   video,
	}, opts), &msg)
	return
}

func (tg *Telegram) SendAnimation(chat Chat, animation interface{}, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("sendAnimation", createPayload(fields{
		"chat_id":   chat.ID,
		"animation": animation,
	}, opts), &msg)
	return
}

func (tg *Telegram) SendVoice(chat Chat, voice interface{}, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("sendVoice", createPayload(fields{
		"chat_id": chat.ID,
		"voice":   voice,
	}, opts), &msg)
	return
}

func (tg *Telegram) SendVideoNote(chat Chat, videoNote interface{}, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("sendVideoNote", createPayload(fields{
		"chat_id":    chat.ID,
		"video_note": videoNote,
	}, opts), &msg)
	return
}

// // mediaGroup - array of audios, documents, photos, videos and links to the media
// func (tg *Telegram) SendMediaGroup(chat Chat, mediaGroup []interface{}, opts ...Option) (msg Message, err error) {
// 	if len(mediaGroup) < 2 || len(mediaGroup) > 10 {
// 		err = errors.New("mediaGroup size is out of range")
// 		return nil, err
// 	}
// 	err = tg.client.Call("sendLocation", createPayload(fields{
// 		"chat_id": chat.ID,
// 		"media":   mediaGroup,
// 	}, opts), msg)
// 	return
// }

func (tg *Telegram) SendLocation(chat Chat, latitude, longitude float32, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("sendLocation", createPayload(fields{
		"chat_id":   chat.ID,
		"latitude":  latitude,
		"longitude": longitude,
	}, opts), &msg)
	return
}

func (tg *Telegram) EditMessageLiveLocation(msg Message, latitude, longitude float32, opts ...Option) (editedMsg Message, err error) {
	if msg.Chat.ID == 0 {
		opts = append(opts, Option{"inline_message_id", msg.MessageID})
	} else {
		opts = append(opts, Option{"chat_id", msg.Chat.ID}, Option{"message_id", msg.MessageID})
	}

	err = tg.client.Call("editMessageLiveLocation", createPayload(fields{
		"latitude":  latitude,
		"longitude": longitude,
	}, opts), &editedMsg)
	return
}

func (tg *Telegram) StopMessageLiveLocation(msg Message, opts ...Option) (editedMsg Message, err error) {
	err = tg.client.Call("stopMessageLiveLocation", createPayload(fields{
		"chat_id":    msg.Chat.ID,
		"message_id": msg.MessageID,
	}, opts), &editedMsg)
	return
}

func (tg *Telegram) SendVenue(chat Chat, latitude, longitude float32, title, address string, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("sendVenue", createPayload(fields{
		"chat_id":   chat.ID,
		"latitude":  latitude,
		"longitude": longitude,
		"title":     title,
		"address":   address,
	}, opts), &msg)
	return
}

func (tg *Telegram) SendContact(chat Chat, phoneNumber, firstName string, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("sendContact", createPayload(fields{
		"chat_id":      chat.ID,
		"phone_number": phoneNumber,
		"first_name":   firstName,
	}, opts), &msg)
	return
}

func (tg *Telegram) SendPoll(chat Chat, question string, variants []string, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("sendPoll", createPayload(fields{
		"chat_id":  chat.ID,
		"question": question,
		"options":  variants,
	}, opts), &msg)
	return
}

func (tg *Telegram) SendDice(chat Chat, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("sendDice", createPayload(fields{
		"chat_id": chat.ID,
	}, opts), &msg)
	return
}

func (tg *Telegram) SendChatAction(chat Chat, action string) (res bool, err error) {
	err = tg.client.Call("sendChatAction", createPayload(fields{
		"chat_id": chat.ID,
		"action":  action,
	}, []Option{}), &res)
	return
}

func (tg *Telegram) GetUserProfilePhotos(user User, opts ...Option) (photos UserProfilePhotos, err error) {
	err = tg.client.Call("getUserProfilePhotos", createPayload(fields{
		"user_id": user.ID,
	}, opts), &photos)
	return
}

func (tg *Telegram) GetFile(fileID string, opts ...Option) (file File, err error) {
	err = tg.client.Call("getFile", createPayload(fields{
		"file_id": fileID,
	}, opts), &file)
	return
}

func (tg *Telegram) AnswerCallbackQuery(cbq CallbackQuery, opts ...Option) (answered bool, err error) {
	err = tg.client.Call("answerCallbackQuery", createPayload(fields{
		"callback_query_id": cbq.ID,
	}, opts), &answered)
	return
}

func (tg *Telegram) SetMyCommands(commands []BotCommand) (setted bool, err error) {
	err = tg.client.Call("setMyCommands", createPayload(fields{
		"commands": commands,
	}, []Option{}), &setted)
	return
}

func (tg *Telegram) GetMyCommands() (commands []BotCommand, err error) {
	err = tg.client.Call("getMyCommands", createPayload(fields{}, []Option{}), &commands)
	return
}
