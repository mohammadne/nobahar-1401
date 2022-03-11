package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Auth struct {
}

func (a Auth) Register(r fiber.Router) {
	r.Post("/auth/signup", a.signup)
	r.Post("/auth/login", a.login)
}

func (a Auth) signup(ctx *fiber.Ctx) error {
	payload := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}

	return ctx.JSON(payload)
}

func (a Auth) login(ctx *fiber.Ctx) error {
	payload := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}

	return ctx.JSON(payload)
}
