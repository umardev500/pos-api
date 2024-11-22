package pkg

func ValidationErrorResponse(fields []ValidationErrorMeta) Response {
	refCode := LogError(ErrValidation)

	return Response{
		Message: "Validation errors occurred.",
		Error: &ErrorDetail{
			Code:    ErrCodeNameValidation,
			Message: "Validation error",
			Errors:  fields,
		},
		RefCode: refCode,
	}
}
