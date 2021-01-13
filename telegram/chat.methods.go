package telegram

func (tg *Telegram) KickChatMember(chat Chat, user User, opts ...Option) (res bool, err error) {
	err = tg.client.Call("kickChatMember", createPayload(fields{
		"chat_id": chat.ID,
		"user_id": user.ID,
	}, opts), &res)
	return
}

func (tg *Telegram) UnbanChatMember(chat Chat, user User, opts ...Option) (res bool, err error) {
	err = tg.client.Call("unbanChatMember", createPayload(fields{
		"chat_id": chat.ID,
		"user_id": user.ID,
	}, opts), &res)
	return
}

func (tg *Telegram) RestrictChatMember(chat Chat, user User, permissions ChatPermissions, opts ...Option) (res bool, err error) {
	err = tg.client.Call("restrictChatMember", createPayload(fields{
		"chat_id":     chat.ID,
		"user_id":     user.ID,
		"permissions": permissions,
	}, opts), &res)
	return
}

func (tg *Telegram) PromoteChatMember(chat Chat, user Chat, opts ...Option) (res bool, err error) {
	err = tg.client.Call("promoteChatMember", createPayload(fields{
		"chat_id": chat.ID,
		"user_id": user.ID,
	}, opts), &res)
	return
}

func (tg *Telegram) SetChatAdministratorCustomTitle(chat Chat, user User, customTitle string, opts ...Option) (res bool, err error) {
	err = tg.client.Call("setChatAdministratorCustomTitle", createPayload(fields{
		"chat_id":      chat.ID,
		"user_id":      user.ID,
		"custom_title": customTitle,
	}, opts), &res)
	return
}

func (tg *Telegram) SetChatPermissions(chat Chat, permissions ChatPermissions, opts ...Option) (res bool, err error) {
	err = tg.client.Call("setChatPermissions", createPayload(fields{
		"chat_id":     chat.ID,
		"permissions": permissions,
	}, opts), &res)
	return
}

func (tg *Telegram) ExportChatInviteLink(chat Chat, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("exportChatInviteLink", createPayload(fields{
		"chat_id": chat.ID,
	}, opts), msg)
	return
}

func (tg *Telegram) SetChatPhoto(chat Chat, photo interface{}, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("setChatPhoto", createPayload(fields{
		"chat_id": chat.ID,
		"photo":   photo,
	}, opts), msg)
	return
}

func (tg *Telegram) DeleteChatPhoto(chat Chat, opts ...Option) (msg Message, err error) {
	err = tg.client.Call("deleteChatPhoto", createPayload(fields{
		"chat_id": chat.ID,
	}, opts), msg)
	return
}

//TODO: test, description is optional
func (tg *Telegram) SetChatTitle(chat Chat, title string) (setted bool, err error) {
	err = tg.client.Call("setChatDescription", createPayload(fields{
		"chat_id": chat.ID,
		"title":   title,
	}, []Option{}), &setted)
	return
}

//TODO: test, description is optional
func (tg *Telegram) SetChatDescription(chat Chat, description string) (setted bool, err error) {
	err = tg.client.Call("setChatDescription", createPayload(fields{
		"chat_id":     chat.ID,
		"description": description,
	}, []Option{}), &setted)
	return
}

func (tg *Telegram) PinChatMessage(chat Chat, msg Message, opts ...Option) (pinned bool, err error) {
	err = tg.client.Call("pinChatMessage", createPayload(fields{
		"chat_id":    chat.ID,
		"message_id": msg.MessageID,
	}, opts), &pinned)
	return
}

func (tg *Telegram) UnpinChatMessage(chat Chat, msg Message) (unpinned bool, err error) {
	err = tg.client.Call("unpinChatMessage", createPayload(fields{
		"chat_id":    chat.ID,
		"message_id": msg.MessageID,
	}, []Option{}), &unpinned)
	return
}

func (tg *Telegram) UnpinAllChatMessages(chat Chat) (unpinned bool, err error) {
	err = tg.client.Call("unpinAllChatMessages", createPayload(fields{
		"chat_id": chat.ID,
	}, []Option{}), &unpinned)
	return
}

func (tg *Telegram) LeaveChat(chat Chat) (leavedChat Chat, err error) {
	err = tg.client.Call("leaveChat", createPayload(fields{
		"chat_id": chat.ID,
	}, []Option{}), &leavedChat)
	return
}

func (tg *Telegram) GetChat(chat Chat) (gettedChat Chat, err error) {
	err = tg.client.Call("getChat", createPayload(fields{
		"chat_id": chat.ID,
	}, []Option{}), &gettedChat)
	return
}

func (tg *Telegram) GetChatAdministrators(chat Chat) (admins []ChatMember, err error) {
	err = tg.client.Call("getChatAdministrators", createPayload(fields{
		"chat_id": chat.ID,
	}, []Option{}), &admins)
	return
}

func (tg *Telegram) GetChatMembersCount(chat Chat) (count int, err error) {
	err = tg.client.Call("getChatMembersCount", createPayload(fields{
		"chat_id": chat.ID,
	}, []Option{}), &count)
	return
}

func (tg *Telegram) GetChatMember(chat Chat, user User) (member ChatMember, err error) {
	err = tg.client.Call("getChatMember", createPayload(fields{
		"chat_id": chat.ID,
		"user_id": user.ID,
	}, []Option{}), &member)
	return
}

func (tg *Telegram) SetChatStickerSet(chat Chat, stickerSetName string) (setted bool, err error) {
	err = tg.client.Call("setChatStickerSet", createPayload(fields{
		"chat_id":          chat.ID,
		"sticker_set_name": stickerSetName,
	}, []Option{}), &setted)
	return
}

func (tg *Telegram) DeleteChatStickerSet(chat Chat) (deleted bool, err error) {
	err = tg.client.Call("deleteChatStickerSet", createPayload(fields{
		"chat_id": chat.ID,
	}, []Option{}), &deleted)
	return
}
