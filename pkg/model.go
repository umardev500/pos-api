// Package pkg contains common models used across the application.
package pkg

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Error and Validation Models

// ErrorDetail represents details of an error.
type ErrorDetail struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

// ValidationErrorMeta represents metadata for validation errors.
type ValidationErrorMeta struct {
	Code    string      `json:"code"`
	Type    string      `json:"type"`
	Field   string      `json:"field"`
	Minimum int         `json:"minimum,omitempty"`
	Maximum int         `json:"maximum,omitempty"`
	Exact   int         `json:"exact,omitempty"`
	Items   interface{} `json:"items,omitempty"`
	Message string      `json:"message"`
}

// Pagination and Sorting Models

// PaginationResponseMeta contains metadata for paginated responses.
type PaginationResponseMeta struct {
	CurrentPage int64 `json:"current_page"`
	PerPage     int64 `json:"per_page"`
	TotalPages  int64 `json:"total_pages"`
	TotalItems  int64 `json:"total_items"`
}

// DateRange represents a range of dates.
type DateRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// Sort defines sorting parameters.
type Sort struct {
	SortBy string   `json:"sort_by"`
	Sort   SortMode `json:"sort"`
}

// Pagination defines pagination parameters.
type Pagination struct {
	Page    int64 `json:"page"`
	PerPage int64 `json:"per_page"`
	Offset  int64 `json:"-"`
}

// FindRequest represents a request for finding resources with filters, pagination, and sorting.
type FindRequest struct {
	Filters    interface{} `json:"filters"`
	Pagination *Pagination `json:"pagination"`
	Sort       *Sort       `json:"sort"`
	Search     *string     `json:"search"`
}

// Response Model

// Response represents a standard API response.
type Response struct {
	StatusCode int                     `json:"-"`
	Success    bool                    `json:"success"`
	Message    string                  `json:"message,omitempty"`
	Data       interface{}             `json:"data,omitempty"`
	Pagination *PaginationResponseMeta `json:"pagination,omitempty"`
	Error      *ErrorDetail            `json:"error,omitempty"`
	RefCode    string                  `json:"ref_code,omitempty"`
}

type TimeModel struct {
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type IdsModel struct {
	IDs []uuid.UUID `json:"ids"`
}

func (i *IdsModel) Validate() error {
	return nil
}
