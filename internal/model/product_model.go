package model

import (
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

// Product represents a product with its details and pricing.
type Product struct {
	ID          string           `json:"id" gorm:"column:id"`
	Name        string           `json:"name" gorm:"column:name"`
	Description string           `json:"description" gorm:"column:description"`
	Pricings    []ProductPricing `json:"pricings" gorm:"foreignKey:product_id"`
	pkg.TimeModel
}

// TableName sets the insert table name for this struct type.
func (p *Product) TableName() string {
	return "products"
}
