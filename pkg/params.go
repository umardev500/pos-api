package pkg

import "github.com/gofiber/fiber/v2"

func GetPaginationParams(c *fiber.Ctx) Pagination {
	page := c.QueryInt("page", 1)
	perPage := c.QueryInt("per_page", 10)
	return Pagination{
		Page:    int64(page),
		PerPage: int64(perPage),
		Offset:  int64((page - 1) * perPage),
	}
}

func GetSortParams(c *fiber.Ctx) Sort {
	sortBy := c.Query("sort_by")
	sort := c.Query("sort")

	// Validate sort mode
	if _, ok := AllowedSortdModes[sort]; !ok {
		sort = "asc"
	}

	return Sort{
		SortBy: sortBy,
		Sort:   SortMode(sort),
	}
}
