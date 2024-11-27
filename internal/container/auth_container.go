package container

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/handler"
	"github.com/umardev500/pos-api/internal/repository"
	"github.com/umardev500/pos-api/internal/service"
	"github.com/umardev500/pos-api/pkg"
)

type authContainer struct {
	hndlr contract.AuthHandler
}

func NewAuthContainer(db *pkg.GormDB, v pkg.Validator) pkg.Container {
	authRepo := repository.NewAuthRepository(db)
	authSrv := service.NewAuthService(authRepo, v)
	authHndlr := handler.NewAuthHandler(authSrv)

	return &authContainer{hndlr: authHndlr}
}

func (a *authContainer) HandleApi(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/login", a.hndlr.Login)
}

func (a *authContainer) HandleWeb(router fiber.Router) {}
