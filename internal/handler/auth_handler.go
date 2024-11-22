package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/model"
)

type authHandler struct {
	authService contract.AuthService
}

func NewAuthHandler(authService contract.AuthService) contract.AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (a *authHandler) Login(c *fiber.Ctx) error {
	var payload model.LoginRequest
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	resp := a.authService.Login(c.Context(), &payload)

	return c.JSON(resp)
}
