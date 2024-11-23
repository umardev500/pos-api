package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-api/internal/contract"
)

type userHandler struct {
	userService contract.UserService
}

func NewUserHandler(us contract.UserService) contract.UserHandler {
	return &userHandler{
		userService: us,
	}
}

func (u *userHandler) HandleGetCurrentUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userId := c.Locals("user_id")
	resp := u.userService.FindUserByID(ctx, userId.(string))

	return c.Status(resp.StatusCode).JSON(resp)
}
