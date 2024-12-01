package service

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
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

func (p *productService) FindAllProducts(ctx context.Context, params pkg.FindRequest) pkg.Response {
	products, total, err := p.repo.FindAllProducts(ctx, params)
	if err != nil {
		return pkg.InternalErrorResponse(err)
	}

	// Validate status
	if err := params.Filters.(*model.ProductFilter).Status.Validate(); err != nil {
		return pkg.BadRequestResponse(fmt.Errorf("invalid status: %w", err))
	}

	return pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Resources found successfully",
		Data:       products,
		Pagination: pkg.ParsePaginationInfo(total, params),
	}
}
