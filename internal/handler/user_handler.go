package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/pkg"
)

type userHandler struct {
	userService contract.UserService
}

func NewUserHandler(us contract.UserService) contract.UserHandler {
	return &userHandler{
		userService: us,
	}
}

func (u *userHandler) HandleGetAllUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pagination := pkg.GetPaginationParams(c)
	sort := pkg.GetSortParams(c)

	params := pkg.FindRequest{
		Pagination: &pagination,
		Sort:       &sort,
	}

	resp := u.userService.FindAllUsers(ctx, params)
	return c.Status(resp.StatusCode).JSON(resp)
}

func (u *userHandler) HandleGetCurrentUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userId := c.Locals("user_id")
	resp := u.userService.FindUserByID(ctx, userId.(string))

	return c.Status(resp.StatusCode).JSON(resp)
}
