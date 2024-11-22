package pkg

import "github.com/gofiber/fiber/v2"

func ValidationErrorResponse(fields []ValidationErrorMeta) Response {
	refCode := LogError(ErrValidation)

	return Response{
		StatusCode: fiber.StatusUnprocessableEntity,
		Message:    "Validation errors occurred.",
		Error: &ErrorDetail{
			Code:    ErrCodeNameValidation,
			Message: "Validation error",
			Errors:  fields,
		},
		RefCode: refCode,
	}
}

func InternalErrorResponse(err error) Response {
	refCode := LogError(err)

	return Response{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "Internal server error.",
		RefCode:    refCode,
	}
}
