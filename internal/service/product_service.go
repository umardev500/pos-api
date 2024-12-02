package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
)

type productService struct {
	repo contract.ProductRepository
	v    pkg.Validator
}

func NewProductService(repo contract.ProductRepository, v pkg.Validator) contract.ProductService {
	return &productService{
		repo: repo,
		v:    v,
	}
}

func (p *productService) DeleteProductById(ctx context.Context, id string) pkg.Response {
	uid, err := uuid.Parse(id)
	if err != nil {
		return pkg.BadRequestResponse(err)
	}

	err = p.repo.DeleteProductById(ctx, uid)
	if err != nil {
		return pkg.InternalErrorResponse(err)
	}

	return pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Resource deleted successfully",
	}
}

func (p *productService) FindAllProducts(ctx context.Context, params pkg.FindRequest) pkg.Response {
	filters := params.Filters.(*model.ProductFilter)
	err := filters.Validate()
	if err != nil {
		return pkg.BadRequestResponse(err)
	}

	products, total, err := p.repo.FindAllProducts(ctx, params)
	if err != nil {
		return pkg.InternalErrorResponse(err)
	}

	return pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Resources found successfully",
		Data:       products,
		Pagination: pkg.ParsePaginationInfo(total, params),
	}
}
