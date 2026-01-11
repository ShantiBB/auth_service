package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	"booking/pkg/utils/consts"
)

func CustomValidationError(err validator.FieldError) string {
	value := err.Value()
	param := err.Param()

	switch err.Tag() {
	case "required":
		return consts.FieldRequired
	case "slug_format":
		return consts.FieldSlug
	case "min":
		return fmt.Sprintf(consts.FieldMin, param)
	case "max":
		return fmt.Sprintf(consts.FieldMax, param)
	case "gt", "decimal_gt":
		return fmt.Sprintf(consts.FieldGt, param, value)
	case "gte":
		return fmt.Sprintf(consts.FieldGte, param, value)
	case "lt", "decimal_lt":
		return fmt.Sprintf(consts.FieldLt, param, value)
	case "lte":
		return fmt.Sprintf(consts.FieldLte, param, value)
	case "uuid":
		return consts.FieldUUID
	case "datetime":
		return fmt.Sprintf(consts.FieldDatetime, param)
	default:
		return fmt.Sprintf(consts.FieldInvalid, value)
	}
}
