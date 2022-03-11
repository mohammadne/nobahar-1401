package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadne/nobahar-1401/internal/db"
	"github.com/mohammadne/nobahar-1401/internal/http/request"
	"github.com/mohammadne/nobahar-1401/internal/http/response"
	"github.com/mohammadne/nobahar-1401/internal/jwt"
	"github.com/mohammadne/nobahar-1401/internal/models"
)

type Auth struct {
	DB  db.DB
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

	user, err := a.DB.GetUser(request.Email)
	if err != nil {
		return err
	} else if user.ID != 0 {
		return errors.New("user exists")
	}

	user = &models.User{Name: request.Name, Email: request.Email}
	if err := a.DB.CreateUser(user); err != nil {
		return err
	}

	token, err := a.JWT.CreateTokenString(user.ID, user.Email)
	if err != nil {
		return err
	}

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
