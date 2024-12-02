package service

import (
	"context"

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

func (p *productService) SoftDeleteProducts(ctx context.Context, req *pkg.IdsModel) pkg.Response {
	err := req.Validate()
	if err != nil {
		return pkg.BadRequestResponse(err)
	}

	err = p.repo.SoftDeleteProducts(ctx, req.IDs)
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

func (p *productService) RestoreDeletedProducts(ctx context.Context, req *pkg.IdsModel) pkg.Response {
	err := req.Validate()
	if err != nil {
		return pkg.BadRequestResponse(err)
	}

	err = p.repo.RestoreDeletedProducts(ctx, req.IDs)
	if err != nil {
		return pkg.InternalErrorResponse(err)
	}

	return pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Resources restored successfully",
	}
}
