package validation

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"

	"auth_service/package/utils/errs"
)

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return msgRequired
	case "email":
		return msgEmail
	case "min":
		return fmt.Sprintf(msgMin, err.Param())
	default:
		return msgDefault
	}
}

func formatValidationErrors(validationErrors validator.ValidationErrors) map[string]string {
	errorMessages := make(map[string]string)

	for _, err := range validationErrors {
		errorMessages[err.Field()] = getErrorMessage(err)
	}

	return errorMessages
}

func CheckErrors(v interface{}) *errs.ValidateError {
	validate := validator.New()
	if err := validate.Struct(v); err != nil {
		var validateErr validator.ValidationErrors
		errors.As(err, &validateErr)

		errorsResp := errs.ValidateError{
			Errors: formatValidationErrors(validateErr),
		}
		return &errorsResp
	}

	return nil
}
