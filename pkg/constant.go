// Package pkg contains common constants and types used across the application.
package pkg

import "fmt"

// ContextKey is a custom type for context keys.
type ContextKey int

// Context keys used in the application.
const (
	TrxKey ContextKey = iota
)

// Error codes used in the application.
const (
	ErrCodeNameValidation string = "VALIDATION_ERROR"
)

// Common errors used in the application.
var (
	ErrValidation error = fmt.Errorf("validation error")
)

// Logic represents logical operators.
type Logic string

// Logical operators.
const (
	LogicAnd Logic = "and"
	LogicOr  Logic = "or"
)

// Operator represents comparison operators.
type Operator string

// Comparison operators.
const (
	OperatorEq   Operator = "="
	OperatorNe   Operator = "!="
	OperatorGt   Operator = ">"
	OperatorGte  Operator = ">="
	OperatorLt   Operator = "<"
	OperatorLte  Operator = "<="
	OperatorLike Operator = "like"
)

type SortMode string

// SortMode represents sort modes.
const (
	SortModeAsc  SortMode = "asc"
	SortModeDesc SortMode = "desc"
)

var AllowedSortdModes = map[string]SortMode{
	"asc":  SortModeAsc,
	"desc": SortModeDesc,
}
