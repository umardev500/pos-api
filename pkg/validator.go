package pkg

import (
	"reflect"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type validatorStruct struct {
	validate *validator.Validate
}

func NewValidator() Validator {
	return &validatorStruct{
		validate: validator.New(),
	}
}

func (v *validatorStruct) GetValidator() *validator.Validate {
	return v.validate
}

func (v *validatorStruct) Struct(obj interface{}) (fields []ValidationErrorMeta, err error) {
	err = v.validate.Struct(obj)
	if err == nil {
		return nil, nil
	}

	val := reflect.ValueOf(obj).Elem()

	fields = make([]ValidationErrorMeta, 0)

	for _, fieldErr := range err.(validator.ValidationErrors) {
		fieldName := fieldErr.Field()
		field, ok := val.Type().FieldByName(fieldName)
		if !ok {
			continue
		}

		validationField := v.GetDetail(fieldErr, field)
		if validationField == nil {
			continue
		}

		fields = append(fields, *validationField)
	}

	return fields, ErrValidation
}

func (v *validatorStruct) GetDetail(fieldErr validator.FieldError, field reflect.StructField) *ValidationErrorMeta {
	code := fieldErr.Tag()
	path, ok := field.Tag.Lookup("name")
	if !ok {
		path = field.Name
	}
	fieldType := field.Type.Kind().String()

	param := fieldErr.Param()

	var validationField = &ValidationErrorMeta{
		Code:  code,
		Type:  fieldType,
		Field: path,
	}

	switch code {
	case "min":
		validationField.Minimum = func() int {
			i, _ := strconv.Atoi(param)
			return i
		}()
		validationField.Message = "Must be at least " + param
		return validationField
	case "max":
		validationField.Maximum = func() int {
			i, _ := strconv.Atoi(param)
			return i
		}()
		validationField.Message = "Must be at most " + param
		return validationField
	case "required":
		validationField.Message = "This field is required"
		return validationField
	case "email":
		validationField.Message = "Invalid email"
		return validationField
	case "len":
		validationField.Exact = func() int {
			i, _ := strconv.Atoi(param)
			return i
		}()
		validationField.Message = "Must be " + param + " characters long"
		return validationField
	case "oneof":
		validationField.Items = param
		validationField.Message = "Must be one of " + param
		return validationField
	case "gt":
		validationField.Exact = func() int {
			i, _ := strconv.Atoi(param)
			return i
		}()
		validationField.Message = "Must be greater than " + param
		return validationField
	case "gte":
		validationField.Exact = func() int {
			i, _ := strconv.Atoi(param)
			return i
		}()
		validationField.Message = "Must be greater than or equal to " + param
		return validationField
	case "lt":
		validationField.Exact = func() int {
			i, _ := strconv.Atoi(param)
			return i
		}()
		validationField.Message = "Must be less than " + param
		return validationField
	case "lte":
		validationField.Exact = func() int {
			i, _ := strconv.Atoi(param)
			return i
		}()
		validationField.Message = "Must be less than or equal to " + param
		return validationField
	default:
		return validationField
	}
}
