package validation

import (
	"github.com/go-playground/validator/v10"

	"hotel/internal/repository/models"
)

func roomStatusValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	for _, v := range models.RoomStatusValues {
		if string(v) == value {
			return true
		}
	}
	return false
}
