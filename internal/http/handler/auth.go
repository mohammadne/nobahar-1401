package handler

import "github.com/gofiber/fiber/v2"

type Auth struct {
}

func (a Auth) Register(app *fiber.App) {
	app.Post("/api/v1/auth/signup", a.signup)
	app.Post("/api/v1/auth/login", a.login)
}

func (a Auth) signup(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (a Auth) login(ctx *fiber.Ctx) error {
	panic("implement me")
}
