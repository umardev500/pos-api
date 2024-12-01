package pkg

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Router interface {
	Handle()
	HandleApi(router fiber.Router)
	HandleWeb(router fiber.Router)
}

type Container interface {
	HandleApi(router fiber.Router)
	HandleWeb(router fiber.Router)
}

type Validator interface {
	GetValidator() *validator.Validate
	Struct(obj interface{}) ([]ValidationErrorMeta, error)
}

type Validation interface {
	Message() string
}

type Filter interface {
	Validate(value interface{}) error
}
