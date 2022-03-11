package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadne/nobahar-1401/internal/http/request"
	"github.com/mohammadne/nobahar-1401/internal/http/response"
	"github.com/mohammadne/nobahar-1401/internal/jwt"
)

type Auth struct {
	JWT jwt.JWT
}

func (a Auth) Register(r fiber.Router) {
	r.Post("/auth/signup", a.signup)
	r.Post("/auth/login", a.login)
}

func (a Auth) signup(ctx *fiber.Ctx) error {
	request := request.SignupRequest{}

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	token, err := a.JWT.CreateTokenString(1, request.Email)
	if err != nil {
		return err
	}

	id, email, err := a.JWT.ExtractTokenData(token)
	if err != nil {
		panic(err)
		return err
	}
	log.Printf("id:%d, email;%s", id, email)

	return ctx.JSON(response.SignupResponse{
		Token:   token,
		Message: "successfull",
	})
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
