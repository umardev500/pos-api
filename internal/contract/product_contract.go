package contract

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
)

type ProductHandler interface {
	HandleDeleteProductById(ctx *fiber.Ctx) error
	HandleGetAllProducts(ctx *fiber.Ctx) error
	HandleRestoreDeletedProducts(ctx *fiber.Ctx) error
}

type ProductService interface {
	DeleteProductById(ctx context.Context, id string) pkg.Response
	FindAllProducts(ctx context.Context, params pkg.FindRequest) pkg.Response
	RestoreDeletedProducts(ctx context.Context, req *pkg.IdsModel) pkg.Response
}

type ProductRepository interface {
	DeleteProductById(ctx context.Context, id uuid.UUID) error
	FindAllProducts(ctx context.Context, params pkg.FindRequest) ([]model.Product, int64, error)
	RestoreDeletedProducts(ctx context.Context, ids []uuid.UUID) error
}
