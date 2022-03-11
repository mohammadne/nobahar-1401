package response

import "github.com/mohammadne/nobahar-1401/internal/http/model"

type GetChatsResponse struct {
	Chats []model.Chat `json:"chats"`
}

type GetChatMessagesResponse struct {
	Messages []model.Message `json:"messages"`
}
