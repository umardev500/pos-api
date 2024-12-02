package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-api/internal/contract"
	"github.com/umardev500/pos-api/internal/model"
	"github.com/umardev500/pos-api/pkg"
)

type productHandler struct {
	service contract.ProductService
}

func NewProductHandler(service contract.ProductService) contract.ProductHandler {
	return &productHandler{service: service}
}

func (ph *productHandler) HandleDeleteProducts(c *fiber.Ctx) error {
	var payload pkg.IdsModel
	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp := ph.service.SoftDeleteProducts(ctx, &payload)
	return c.Status(resp.StatusCode).JSON(resp)
}

func (ph *productHandler) HandleGetAllProducts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pagination := pkg.GetPaginationParams(c)
	sort := pkg.GetSortParams(c)
	s := c.Query("search")
	status := c.Query("status")
	category := c.Query("category")
	minPrice := c.QueryFloat("min_price", 0)
	maxPrice := c.QueryFloat("max_price", 0)
	archived := c.QueryBool("archived")

	filters := model.ProductFilter{
		Status:   (*model.ProductStatus)(&status),
		Archived: archived,
		Category: &category,
		MinPrice: &minPrice,
		MaxPrice: &maxPrice,
	}

	params := pkg.FindRequest{
		Filters:    &filters,
		Pagination: &pagination,
		Sort:       &sort,
		Search:     &s,
	}

	resp := ph.service.FindAllProducts(ctx, params)
	return c.Status(resp.StatusCode).JSON(resp)
}

func (ph *productHandler) HandleRestoreDeletedProducts(c *fiber.Ctx) error {
	var payload pkg.IdsModel
	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp := ph.service.RestoreDeletedProducts(ctx, &payload)
	return c.Status(resp.StatusCode).JSON(resp)
}
