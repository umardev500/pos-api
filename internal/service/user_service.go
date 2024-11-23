package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/pkg"
)

type userService struct {
	repo contract.UserRepository
	v    pkg.Validator
}

func NewUserService(repo contract.UserRepository, v pkg.Validator) contract.UserService {
	return &userService{
		repo: repo,
		v:    v,
	}
}

func (u *userService) FindUserByID(ctx context.Context, id string) pkg.Response {
	user, err := u.repo.FindUserById(ctx, id)
	if err != nil {
		if err = pgx.ErrNoRows; err != nil {
			return pkg.NotFoundResponse(nil)
		}

		return pkg.InternalErrorResponse(err)
	}

	return pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Resource found successfully",
		Data:       user,
	}
}
