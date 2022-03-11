package http

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadne/nobahar-1401/internal/http/handler"
)

type server struct {
	app    *fiber.App
	config *Config
}

func New(cfg *Config) *server {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	handler.Auth{}.Register(app)
	handler.Chat{}.Register(app)
	handler.Group{}.Register(app)

	return &server{app: app, config: cfg}
}

func (s *server) Serve() {
	baseAddress := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	log.Printf("server started on: %s\n", baseAddress)
	if err := s.app.Listen(baseAddress); err != nil {
		panic("error while starting the server")
	}
}
