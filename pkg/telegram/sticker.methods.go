package telegram

import (
	"encoding/json"

	"github.com/lepeico/gogram/internal/client"
)

func (tg *Telegram) SendSticker(chatID string, sticker interface{}) (json.RawMessage, error) {
	message := client.Payload{}
	message.Set("chat_id", chatID)
	message.Set("sticker", sticker)

	return tg.Call("sendSticker", message)
}

func (tg *Telegram) GetStickerSet(name string) (json.RawMessage, error) {
	message := client.Payload{}
	message.Set("name", name)

	return tg.Call("getStickerSet", message)
}

func (tg *Telegram) UploadStickerFile(userID string, pngSticker interface{}) (json.RawMessage, error) {
	message := client.Payload{}
	message.Set("user_id", userID)
	message.Set("png_sticker", pngSticker)

	return tg.Call("uploadStickerFile", message)
}

//TODO: add tgs_sticker & arrray of emojis
func (tg *Telegram) CreateNewStickerSet(userID, name, title string, pngSticker interface{}, emojis string) (json.RawMessage, error) {
	message := client.Payload{}
	message.Set("user_id", userID)
	message.Set("name", name)
	message.Set("title", title)
	message.Set("png_sticker", pngSticker)
	message.Set("emojis", emojis)

	return tg.Call("createNewStickerSet", message)
}

func (tg *Telegram) AddStickerToSet(userID, name, pngSticker interface{}, emojis string) (json.RawMessage, error) {
	message := client.Payload{}
	message.Set("user_id", userID)
	message.Set("name", name)
	message.Set("png_sticker", pngSticker)
	message.Set("emojis", emojis)

	return tg.Call("addStickerToSet", message)
}

func (tg *Telegram) SetStickerPositionInSet(stickerID, position string) (json.RawMessage, error) {
	message := client.Payload{}
	message.Set("sticker", stickerID)
	message.Set("position", position)

	return tg.Call("setStickerPositionInSet", message)
}

func (tg *Telegram) DeleteStickerFromSet(stickerID string) (json.RawMessage, error) {
	message := client.Payload{}
	message.Set("sticker", stickerID)

	return tg.Call("deleteStickerFromSet", message)
}

func (tg *Telegram) SetStickerSetThumb(name, userID string, thumb interface{}) (json.RawMessage, error) {
	message := client.Payload{}
	message.Set("name", name)
	message.Set("user_id", userID)
	message.Set("thumb", thumb)

	return tg.Call("setStickerSetThumb", message)
}
