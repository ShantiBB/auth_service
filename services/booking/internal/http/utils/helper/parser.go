package helper

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	"booking/internal/http/utils/validation"
	"booking/pkg/utils/consts"
)

func ParseJSON(
	w http.ResponseWriter, r *http.Request,
	v any,
	customErr func(validator.FieldError) string,
) error {
	if err := render.DecodeJSON(r.Body, v); err != nil {
		errMsg := validation.ErrorResp(consts.InvalidJSON)
		SendError(w, r, http.StatusBadRequest, errMsg)
		return err
	}

	if errMsg := validation.CheckErrors(v, customErr); errMsg != nil {
		SendError(w, r, http.StatusBadRequest, errMsg)
		return consts.InvalidJSON
	}

	return nil
}
