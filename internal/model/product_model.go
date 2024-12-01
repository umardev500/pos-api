package model

import (
	"fmt"

	"github.com/umardev500/pos-api/pkg"
)

// ProductPrice represents the name and price of a product.
type ProductPrice struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// ProductUnit represents a unit of measurement for a product.
type ProductUnit struct {
	ID         string  `json:"id,omitempty" gorm:"column:id"`
	Name       string  `json:"name" gorm:"column:name"`
	Multiplier float64 `json:"multiplier" gorm:"column:multiplier"`
}

// TableName sets the insert table name for this struct type.
func (p *ProductUnit) TableName() string {
	return "units"
}

// ProductCustomUnit represents a custom unit of measurement.
type ProductCustomUnit struct {
	ID         string  `json:"id,omitempty" gorm:"column:id"`
	Name       string  `json:"name" gorm:"column:name"`
	Multiplier float64 `json:"multiplier" gorm:"column:multiplier"`
}

// TableName sets the insert table name for this struct type.
func (p *ProductCustomUnit) TableName() string {
	return "custom_units"
}

// ProductPricing represents the pricing details for a product.
type ProductPricing struct {
	ID           string             `json:"id" gorm:"column:id"`
	ProductID    string             `json:"-" gorm:"column:product_id"`
	UnitID       string             `json:"-" gorm:"column:unit_id"`
	CustomUnitID string             `json:"-" gorm:"column:custom_unit_id"`
	TierID       string             `json:"-" gorm:"column:tier_id"`
	Price        float64            `json:"price" gorm:"column:price"`
	Unit         *ProductUnit       `json:"unit,omitempty" gorm:"foreignKey:unit_id"`
	CustomUnit   *ProductCustomUnit `json:"custom_unit,omitempty" gorm:"foreignKey:custom_unit_id"`
}

// TableName sets the insert table name for this struct type.
func (p *ProductPricing) TableName() string {
	return "product_pricings"
}

// ProductStock represents the stock details for a product.
type ProductStock struct {
	ID        string `json:"id" gorm:"column:id"`
	ProductID string `json:"-" gorm:"column:product_id"`
	Quantity  int    `json:"quantity" gorm:"column:quantity"`
	pkg.TimeModel
}

// TableName sets the insert table name for this struct type.
func (p *ProductStock) TableName() string {
	return "product_stocks"
}

// Product represents a product with its details and pricing.
type Product struct {
	ID          string           `json:"id" gorm:"column:id"`
	Name        string           `json:"name" gorm:"column:name"`
	Category    *string          `json:"category" gorm:"column:category_name"`
	Description string           `json:"description" gorm:"column:description"`
	Pricings    []ProductPricing `json:"pricings" gorm:"foreignKey:product_id"`
	Stock       *ProductStock    `json:"stock" gorm:"foreignKey:product_id"`
	pkg.TimeModel
}

// TableName sets the insert table name for this struct type.
func (p *Product) TableName() string {
	return "products"
}

type ProductStatus string

const (
	ProductStatusLowStock ProductStatus = "low_stock"
	ProductStatusInStock  ProductStatus = "in_stock"
	ProductStatusOutStock ProductStatus = "out_stock"
)

func (p ProductStatus) Validate() error {
	// Check if the product status is empty
	// If it is empty, return nil
	if p == "" {
		return nil
	}

	// Check if the product status is valid
	switch p {
	case ProductStatusLowStock, ProductStatusInStock, ProductStatusOutStock:
		return nil
	default:
		return fmt.Errorf("invalid product status: %s", p)
	}
}

type ProductFilter struct {
	Status   *ProductStatus `json:"status"`
	Category *string        `json:"category"`
	MinPrice *float64       `json:"min_price"`
	MaxPrice *float64       `json:"max_price"`
}

func (p *ProductFilter) Validate() error {
	// Validate status
	if err := p.Status.Validate(); err != nil {
		return err
	}

	// Validate price range
	// validate min must be less than max and max must be greater than min
	if (p.MinPrice != nil && *p.MinPrice > 0) && (p.MaxPrice != nil && *p.MaxPrice > 0) {
		if *p.MinPrice > *p.MaxPrice || *p.MaxPrice < *p.MinPrice {
			return fmt.Errorf("invalid price range")
		}
	}

	return nil
}
