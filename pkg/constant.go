package pkg

import "fmt"

type ContextKey int

const (
	TrxKey ContextKey = iota
)

const (
	ErrCodeNameValidation string = "VALIDATION_ERROR"
)

// Error
var (
	ErrValidation error = fmt.Errorf("validation error")
)
