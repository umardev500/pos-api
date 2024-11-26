package contract

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
)

type UserHandler interface {
	HandleGetAllUsers(ctx *fiber.Ctx) error
	HandleGetCurrentUser(ctx *fiber.Ctx) error
}

type UserService interface {
	FindAllUsers(ctx context.Context, params pkg.FindRequest) pkg.Response
	FindUserByID(ctx context.Context, userID string) pkg.Response
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	FindAllUsers(ctx context.Context, params pkg.FindRequest) ([]model.User, int64, error)
	FindUserById(ctx context.Context, id uuid.UUID) (*model.User, error)
}
