package telegram

func (tg *Telegram) SendSticker(chat Chat, sticker interface{}, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("sendSticker", createPayload(fields{
		"chat_id": chat.ID,
		"sticker": sticker,
	}, opts), &msg)
	return
}

func (tg *Telegram) GetStickerSet(name string) (set StickerSet, err error) {
	err = tg.client.Call("getStickerSet", createPayload(fields{
		"name": name,
	}, []Option{}), &set)
	return
}

func (tg *Telegram) UploadStickerFile(user User, sticker interface{}) (file File, err error) {
	err = tg.client.Call("uploadStickerFile", createPayload(fields{
		"user_id":     user.ID,
		"png_sticker": sticker,
	}, []Option{}), &file)
	return
}

//TODO: add tgs_sticker & arrray of emojis
func (tg *Telegram) CreateNewStickerSet(user User, name, title string, sticker interface{}, emojis string, opts ...Option) (created bool, err error) {
	err = tg.client.Call("createNewStickerSet", createPayload(fields{
		"user_id": user.ID,
		"name":    name,
		"title":   title,
		"emojis":  emojis,
	}, opts), &created)
	return
}

func (tg *Telegram) AddStickerToSet(user User, name, sticker interface{}, emojis string, opts ...Option) (added bool, err error) {
	err = tg.client.Call("addStickerToSet", createPayload(fields{
		"user_id": user.ID,
		"name":    name,
		"emojis":  emojis,
	}, opts), &added)
	return
}

func (tg *Telegram) SetStickerPositionInSet(sticker Sticker, position int) (setted bool, err error) {
	err = tg.client.Call("setStickerPositionInSet", createPayload(fields{
		"sticker":  sticker.FileID,
		"position": position,
	}, []Option{}), &setted)
	return
}

func (tg *Telegram) DeleteStickerFromSet(sticker Sticker) (deleted bool, err error) {
	err = tg.client.Call("deleteStickerFromSet", createPayload(fields{
		"sticker": sticker.FileID,
	}, []Option{}), &deleted)
	return
}

func (tg *Telegram) SetStickerSetThumb(name string, user User, thumb interface{}) (setted bool, err error) {
	err = tg.client.Call("setStickerSetThumb", createPayload(fields{
		"name":    name,
		"user_id": user.ID,
		"thumb":   thumb,
	}, []Option{}), &setted)
	return
}
