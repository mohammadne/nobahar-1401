package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadne/nobahar-1401/internal/http/handler"
)

type server struct {
	app *fiber.App
}

func New() *server {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	handler.Auth{}.Register(app)
	handler.Chat{}.Register(app)
	handler.Group{}.Register(app)

	return &server{app}
}

func (s *server) Serve() {
	baseAddress := fmt.Sprintf(":%d", 3000)
	if err := s.app.Listen(baseAddress); err != nil {
		panic("error while starting the server")
	}
}
