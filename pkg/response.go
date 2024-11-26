package pkg

import "github.com/gofiber/fiber/v2"

func BadRequestResponse(err error) Response {
	refCode := LogError(err)

	return Response{
		StatusCode: fiber.StatusBadRequest,
		Message:    "Bad request.",
		RefCode:    refCode,
	}
}

func NotFoundResponse(msg *string) Response {
	if msg == nil {
		defaultMsg := "Resource not found"
		msg = &defaultMsg
	}

	return Response{
		StatusCode: fiber.StatusNotFound,
		Message:    *msg,
	}
}

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
