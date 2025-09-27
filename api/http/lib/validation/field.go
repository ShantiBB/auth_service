package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(validationErrors validator.ValidationErrors) map[string]string {
	errorMessages := make(map[string]string)

	for _, err := range validationErrors {
		errorMessages[err.Field()] = getErrorMessage(err)
	}

	return errorMessages
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "Поле обязательно для заполнения"
	case "email":
		return "Неверный формат email"
	case "min":
		return fmt.Sprintf("Минимальная длина %s символов", err.Param())
	default:
		return "Некорректное значение"
	}
}
