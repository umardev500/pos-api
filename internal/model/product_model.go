package model

import "github.com/umardev500/pos-api/pkg"

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

type Product struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Pricing     []ProductPricing `json:"pricing"`
	pkg.TimeModel
}
