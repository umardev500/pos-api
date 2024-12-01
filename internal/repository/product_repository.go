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

func (p *productRepository) queryByCategory(db *gorm.DB, catName string) {
	db.Where("categories.name = ?", catName)
}

func (p *productRepository) queryByMinPrice(db *gorm.DB, minPrice float64) {
	db.Joins("JOIN product_pricings pp ON pp.product_id = products.id").
		Where("pp.price >= ?", minPrice).
		// Select distinct to prevent duplicate results
		Distinct("products.id", "products.name", "products.description", "products.created_at", "products.updated_at", "products.deleted_at", "categories.name as category_name")
}

func (p *productRepository) queryByMaxPrice(db *gorm.DB, maxPrice float64) {
	db.Joins("JOIN product_pricings pp ON pp.product_id = products.id").
		Where("pp.price <= ?", maxPrice).
		// Select distinct to prevent duplicate results
		Distinct("products.id", "products.name", "products.description", "products.created_at", "products.updated_at", "products.deleted_at", "categories.name as category_name")
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

	// Filter by category
	if filters.Category != nil && *filters.Category != "" {
		p.queryByCategory(result, *filters.Category)
	}

	// Filter by min price
	if filters.MinPrice != nil && *filters.MinPrice > 0 {
		p.queryByMinPrice(result, *filters.MinPrice)
	}

	// Filter by max price
	if filters.MaxPrice != nil && *filters.MaxPrice > 0 {
		p.queryByMaxPrice(result, *filters.MaxPrice)
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

	// Join with categories table
	result.Joins("JOIN categories ON categories.id = products.category_id")

	// Select specific columns
	// this will replaced if we use select distinct
	result.Select("products.id", "products.name", "products.description", "products.category_id", "products.created_at", "products.updated_at", "products.deleted_at", "categories.name as category_name")

	// Apply search filter if provided
	if params.Search != nil && *params.Search != "" {
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

	// Join with categories table in the count query as well
	countQuery.Joins("JOIN categories ON categories.id = products.category_id")

	// Apply the same filters to the count query
	if params.Search != nil && *params.Search != "" {
		countQuery = countQuery.Where("products.name ILIKE ?", "%"+*params.Search+"%")
	}
	p.parseFilter(filters, countQuery)

	// Apply distinct to count query to prevent duplicate counting
	countQuery.Distinct("products.id")

	// Count the total number of distinct products with the applied filters
	countQuery.Count(&count)

	return products, count, nil
}
