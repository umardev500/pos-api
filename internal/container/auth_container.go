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
	hndlr contract.AuthHandler
}

func NewAuthContainer(db *pkg.PGX, v pkg.Validator) pkg.Container {
	authRepo := repository.NewAuthRepository(db)
	authSrv := service.NewAuthService(authRepo, v)
	authHndlr := handler.NewAuthHandler(authSrv)

	return &userContainer{hndlr: authHndlr}
}

func (u *userContainer) HandleApi(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/login", u.hndlr.Login)
}

func (u *userContainer) HandleWeb(router fiber.Router) {}
