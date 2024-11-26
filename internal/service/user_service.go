package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func (u *userService) FindAllUsers(ctx context.Context) pkg.Response {
	users, err := u.repo.FindAllUsers(ctx)
	if err != nil {
		return pkg.InternalErrorResponse(err)
	}

	return pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Resources found successfully",
		Data:       users,
	}
}

func (u *userService) FindUserByID(ctx context.Context, id string) pkg.Response {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return pkg.BadRequestResponse(err)
	}

	user, err := u.repo.FindUserById(ctx, parsedID)
	if err != nil {
		if err == pgx.ErrNoRows {
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
