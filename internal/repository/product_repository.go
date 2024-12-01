package repository

import (
	"context"

	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
	"gorm.io/gorm"
)

type productRepository struct {
	db *pkg.GormDB
}

func NewProductRepository(db *pkg.GormDB) contract.ProductRepository {
	return &productRepository{db: db}
}

func (p *productRepository) queryInstock(db *gorm.DB) {
	db.Joins("JOIN product_stocks ps ON ps.product_id = products.id").
		Where("ps.quantity > ?", 0)
}

func (p *productRepository) queryOutstock(db *gorm.DB) {
	db.Joins("JOIN product_stocks ps ON ps.product_id = products.id").
		Where("ps.quantity = ?", 0)
}

func (p *productRepository) queryLowstock(db *gorm.DB) {
	db.Joins("JOIN product_stocks ps ON ps.product_id = products.id").
		Where("ps.quantity < ps.minimum_quantity")
}

func (p *productRepository) parseFilter(filters *model.ProductFilter, result *gorm.DB) {
	// Filter by status
	if filters.Status != nil {
		switch *filters.Status {
		case model.ProductStatusLowStock:
			p.queryLowstock(result)
		case model.ProductStatusOutStock:
			p.queryOutstock(result)
		case model.ProductStatusInStock:
			p.queryInstock(result)
		}
	}
}

func (p *productRepository) FindAllProducts(ctx context.Context, params pkg.FindRequest) ([]model.Product, int64, error) {
	pagination := params.Pagination
	conn := p.db.GetConn(ctx)
	var products = make([]model.Product, 0)
	var count int64 = 0

	// Prepare the base query with pagination and preloading
	result := conn.Offset(int(pagination.Offset)).Limit(int(pagination.PerPage)).
		Preload("Pricings").
		Preload("Pricings.Unit").
		Preload("Pricings.CustomUnit").
		Preload("Stock")

	// Apply search filter if provided
	if params.Search != nil {
		result = result.Where("products.name ILIKE ?", "%"+*params.Search+"%")
	}

	// Apply status filter if provided
	filters := params.Filters.(*model.ProductFilter)
	p.parseFilter(filters, result)

	// Execute the query to find products
	result.Find(&products)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	// Create a separate query for counting products
	countQuery := conn.Model(&model.Product{})

	// Apply the same filters to the count query
	if params.Search != nil {
		countQuery = countQuery.Where("products.name ILIKE ?", "%"+*params.Search+"%")
	}
	p.parseFilter(filters, countQuery)

	// Count the total number of products with the applied filters
	countQuery.Count(&count)

	return products, count, nil
}
