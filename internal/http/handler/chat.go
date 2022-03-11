package handler

import "github.com/gofiber/fiber/v2"

type Chat struct {
}

func (c Chat) Register(r fiber.Router) {
	r.Post("/chats/:user_id", c.sendMessage)
	r.Get("/chats/:user_id", c.getMessages)
	r.Get("/chats", c.getChats)
}

// sendMessage
// if users are not in the same group or their groups are not connected, returns 400
func (c Chat) sendMessage(ctx *fiber.Ctx) error {
	panic("implement me")
}

// getMessages return all messages between two users
// newest to oldest
// if users didn't chat before or other user doesn't exist, returns 400
func (c Chat) getMessages(ctx *fiber.Ctx) error {
	panic("implement me")
}

// getChats return all chats for user
// newest to oldest
// if no chat exists, return empty array
func (c Chat) getChats(ctx *fiber.Ctx) error {
	panic("implement me")
}
