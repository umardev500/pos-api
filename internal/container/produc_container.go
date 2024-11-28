package container

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/handler"
	"github.com/umardev500/pos-api/internal/repository"
	"github.com/umardev500/pos-api/internal/service"
	"github.com/umardev500/pos-api/pkg"
)

type productContainer struct {
	hndlr contract.ProductHandler
}

func NewProductContainer(db *pkg.GormDB, v pkg.Validator) pkg.Container {
	pRepo := repository.NewProductRepository(db)
	pSrv := service.NewProductService(pRepo, v)
	pHndlr := handler.NewProductHandler(pSrv)

	return &productContainer{
		hndlr: pHndlr,
	}
}

func (c *productContainer) HandleApi(router fiber.Router) {
	product := router.Group("/products")
	product.Use(pkg.CheckAuth())
	product.Get("/", c.hndlr.HandleGetAllProducts)
}

func (c *productContainer) HandleWeb(router fiber.Router) {}
