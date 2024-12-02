package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
	"gorm.io/gorm"
)

// Constants for SQL queries and column names
const (
	joinProductStocks      = "JOIN product_stocks ps ON ps.product_id = products.id"
	joinProductPricings    = "JOIN product_pricings pp ON pp.product_id = products.id"
	joinCategories         = "JOIN categories ON categories.id = products.category_id"
	selectDistinctProducts = "products.id, products.name, products.description, products.created_at, products.updated_at, products.deleted_at, categories.name as category_name"
)

type productRepository struct {
	db *pkg.GormDB
}

func NewProductRepository(db *pkg.GormDB) contract.ProductRepository {
	return &productRepository{db: db}
}

// queryInstock filters products that are in stock
func (p *productRepository) queryInstock(db *gorm.DB) {
	db.Joins(joinProductStocks).
		Where("ps.quantity > ?", 0)
}

// queryOutstock filters products that are out of stock
func (p *productRepository) queryOutstock(db *gorm.DB) {
	db.Joins(joinProductStocks).
		Where("ps.quantity = ?", 0)
}

// queryLowstock filters products that have low stock
func (p *productRepository) queryLowstock(db *gorm.DB) {
	db.Joins(joinProductStocks).
		Where("ps.quantity < ps.minimum_quantity")
}

// queryByCategory filters products by category
func (p *productRepository) queryByCategory(db *gorm.DB, catName string) {
	db.Where("categories.name = ?", catName)
}

// queryWithMinAndMaxPrice filters products by minimum and maximum price
func (p *productRepository) queryWithMinAndMaxPrice(db *gorm.DB, minPrice float64, maxPrice float64) {
	db.Joins(joinProductPricings).
		Where("pp.price >= ? AND pp.price <= ?", minPrice, maxPrice).
		Distinct(selectDistinctProducts)
}

// queryByMinPrice filters products by minimum price
func (p *productRepository) queryByMinPrice(db *gorm.DB, minPrice float64) {
	db.Joins(joinProductPricings).
		Where("pp.price >= ?", minPrice).
		Distinct(selectDistinctProducts)
}

// queryByMaxPrice filters products by maximum price
func (p *productRepository) queryByMaxPrice(db *gorm.DB, maxPrice float64) {
	db.Joins(joinProductPricings).
		Where("pp.price <= ?", maxPrice).
		Distinct(selectDistinctProducts)
}

// parseFilter applies filters to the query
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

	// Filter by min and max price
	isMinAndMaxPriceSet := filters.MinPrice != nil && *filters.MinPrice > 0 && filters.MaxPrice != nil && *filters.MaxPrice > 0
	if isMinAndMaxPriceSet {
		p.queryWithMinAndMaxPrice(result, *filters.MinPrice, *filters.MaxPrice)
	} else {
		// Filter by min price
		if filters.MinPrice != nil && *filters.MinPrice > 0 {
			p.queryByMinPrice(result, *filters.MinPrice)
		}

		// Filter by max price
		if filters.MaxPrice != nil && *filters.MaxPrice > 0 {
			p.queryByMaxPrice(result, *filters.MaxPrice)
		}
	}
}

func (p *productRepository) DeleteProductById(ctx context.Context, id uuid.UUID) error {
	conn := p.db.GetConn(ctx)
	return conn.Delete(&model.Product{}, id).Error
}

// FindAllProducts retrieves all products with pagination and filtering
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
	result.Joins(joinCategories)

	// Select specific columns
	result.Select("products.id", "products.name", "products.description", "products.category_id", "products.created_at", "products.updated_at", "products.deleted_at", "categories.name as category_name")

	// Apply search filter if provided
	if params.Search != nil && *params.Search != "" {
		result = result.Where("products.name ILIKE ?", "%"+*params.Search+"%")
	}

	// Apply filters
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
	countQuery.Joins(joinCategories)

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
