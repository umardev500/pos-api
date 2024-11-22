package contract

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
)

type AuthRepository interface {
	GetUserByUsernameOrEmail(ctx context.Context, username string) (*model.User, error)
}

type AuthService interface {
	Login(ctx context.Context, req *model.LoginRequest) pkg.Response
}

type AuthHandler interface {
	Login(ctx *fiber.Ctx) error
}
