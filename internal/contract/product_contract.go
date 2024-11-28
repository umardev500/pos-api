package contract

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
)

type ProductHandler interface {
	HandleGetAllProducts(ctx *fiber.Ctx) error
}

type ProductService interface {
	FindAllProducts(ctx context.Context, params pkg.FindRequest) pkg.Response
}

type ProductRepository interface {
	FindAllProducts(ctx context.Context, params pkg.FindRequest) ([]model.Product, int64, error)
}
