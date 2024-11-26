package container

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/handler"
	"github.com/umardev500/pos-api/internal/repository"
	"github.com/umardev500/pos-api/internal/service"
	"github.com/umardev500/pos-api/pkg"
)

type userContainer struct {
	hndlr contract.UserHandler
}

func NewUserContainer(db *pkg.PGX, v pkg.Validator) pkg.Container {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, v)
	handlr := handler.NewUserHandler(userService)

	return &userContainer{
		hndlr: handlr,
	}
}

func (u *userContainer) HandleApi(router fiber.Router) {
	user := router.Group("/user")
	user.Use(pkg.CheckAuth())

	user.Get("/me", u.hndlr.HandleGetCurrentUser)
	user.Get("/", u.hndlr.HandleGetAllUsers)
}

func (u *userContainer) HandleWeb(router fiber.Router) {}
