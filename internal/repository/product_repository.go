package repository

import (
	"context"

	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
)

type productRepository struct {
	db *pkg.GormDB
}

func NewProductRepository(db *pkg.GormDB) contract.ProductRepository {
	return &productRepository{db: db}
}

func (p *productRepository) FindAllProducts(ctx context.Context, params pkg.FindRequest) ([]model.Product, int64, error) {
	pagination := params.Pagination
	conn := p.db.GetConn(ctx)
	var products = make([]model.Product, 0)
	var count int64 = 0

	result := conn.Offset(int(pagination.Offset)).Limit(int(pagination.PerPage)).Find(&products)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	conn.Model(&model.Product{}).Count(&count)

	return products, count, nil
}
