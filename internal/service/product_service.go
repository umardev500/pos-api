package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
)

// productService implements the contract.ProductService interface
// and provides methods for product-related operations.
type productService struct {
	repo contract.ProductRepository
	v    pkg.Validator
}

// NewProductService creates a new instance of productService.
func NewProductService(repo contract.ProductRepository, v pkg.Validator) contract.ProductService {
	return &productService{
		repo: repo,
		v:    v,
	}
}

// SoftDeleteProducts soft deletes the products with the given IDs.
func (p *productService) SoftDeleteProducts(ctx context.Context, req *pkg.IdsModel) pkg.Response {
	// Validate the request
	if err := req.Validate(); err != nil {
		return pkg.BadRequestResponse(err)
	}

	// Soft delete products from the repository
	rowsAffected, err := p.repo.SoftDeleteProducts(ctx, req.IDs)
	if err != nil {
		// Return internal error response if repository operation fails
		return pkg.InternalErrorResponse(err)
	}

	// Return successful response with rows affected
	return pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Resource deleted successfully",
		Data:       rowsAffected,
	}
}

// FindAllProducts retrieves all products based on the provided filters.
func (p *productService) FindAllProducts(ctx context.Context, params pkg.FindRequest) pkg.Response {
	// Extract product filters from the request
	filters := params.Filters.(*model.ProductFilter)

	// Validate the filters
	if err := filters.Validate(); err != nil {
		return pkg.BadRequestResponse(err)
	}

	// Retrieve products from the repository
	products, total, err := p.repo.FindAllProducts(ctx, params)
	if err != nil {
		// Return internal error response if repository operation fails
		return pkg.InternalErrorResponse(err)
	}

	// Return successful response with products and pagination info
	return pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Resources found successfully",
		Data:       products,
		Pagination: pkg.ParsePaginationInfo(total, params),
	}
}

// RestoreDeletedProducts restores the soft-deleted products with the given IDs.
func (p *productService) RestoreDeletedProducts(ctx context.Context, req *pkg.IdsModel) pkg.Response {
	// Validate the request
	if err := req.Validate(); err != nil {
		return pkg.BadRequestResponse(err)
	}

	// Restore products from the repository
	rowsAffected, err := p.repo.RestoreDeletedProducts(ctx, req.IDs)
	if err != nil {
		// Return internal error response if repository operation fails
		return pkg.InternalErrorResponse(err)
	}

	// Return successful response with rows affected
	return pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Resources restored successfully",
		Data:       rowsAffected,
	}
}
