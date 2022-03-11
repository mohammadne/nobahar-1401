package response

import "github.com/mohammadne/nobahar-1401/internal/models"

type GetChatsResponse struct {
	Chats []models.Chat `json:"chats"`
}

type GetChatMessagesResponse struct {
	Messages []models.Message `json:"messages"`
}
