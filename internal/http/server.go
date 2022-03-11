package http

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadne/nobahar-1401/internal/http/handler"
	"github.com/mohammadne/nobahar-1401/internal/jwt"
)

type server struct {
	app    *fiber.App
	config *Config
}

func New(cfg *Config, jwt jwt.JWT) *server {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	router := app.Group("/api/v1")
	handler.Auth{JWT: jwt}.Register(router)
	handler.Chat{}.Register(router)
	handler.Group{}.Register(router)

	return &server{app: app, config: cfg}
}

func (s *server) Serve() {
	baseAddress := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	log.Printf("server started on: %s\n", baseAddress)
	if err := s.app.Listen(baseAddress); err != nil {
		panic("error while starting the server")
	}
}
