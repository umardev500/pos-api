package contract

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
)

type UserHandler interface {
	HandleGetCurrentUser(ctx *fiber.Ctx) error
}

type UserService interface {
	FindUserByID(ctx context.Context, userID string) pkg.Response
}

type UserRepository interface {
	FindUserById(ctx context.Context, id string) (*model.User, error)
}
