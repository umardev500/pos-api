package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/umardev500/pos-api/pkg"
)

type ProductPrice struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductUnit struct {
	Name       string  `json:"name"`
	Multiplier float64 `json:"multiplier"`
}

type ProductPricing struct {
	Unit   ProductUnit    `json:"unit"`
	Prices []ProductPrice `json:"prices"`
}

type productPricingList []ProductPricing

func (p *productPricingList) Value() (driver.Value, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *productPricingList) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSON value")
	}

	return json.Unmarshal(bytes, &p)
}

type Product struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Pricing     productPricingList `json:"pricing" gorm:"type:jsonb"`
	pkg.TimeModel
}

func (p *Product) TableName() string {
	return "products"
}
