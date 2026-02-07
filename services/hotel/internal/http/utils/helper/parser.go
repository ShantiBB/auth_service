package helper

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/ShantiBB/fukuro-reserve/services/hotel/internal/http/dto/request"
	"github.com/ShantiBB/fukuro-reserve/services/hotel/internal/http/utils/mapper"
	"github.com/ShantiBB/fukuro-reserve/services/hotel/internal/http/utils/validation"
	"github.com/ShantiBB/fukuro-reserve/services/hotel/internal/repository/models"
	"github.com/ShantiBB/fukuro-reserve/services/hotel/pkg/lib/utils/consts"
)

func ParseJSON(
	w http.ResponseWriter, r *http.Request,
	v any,
	customErr func(validator.FieldError) string,
) error {
	if err := render.DecodeJSON(r.Body, v); err != nil {
		errMsg := validation.ErrorResp(consts.ErrInvalidJSON)
		SendError(w, r, http.StatusBadRequest, errMsg)
		return err
	}

	if errMsg := validation.CheckErrors(v, customErr); errMsg != nil {
		SendError(w, r, http.StatusBadRequest, errMsg)
		return consts.ErrInvalidJSON
	}

	return nil
}

func ParseHotelPathParams(r *http.Request) (models.HotelRef, *validation.ValidateError) {
	pathParams := request.HotelPathParams{
		CountryCode: chi.URLParam(r, "countryCode"),
		CitySlug:    chi.URLParam(r, "citySlug"),
		HotelSlug:   chi.URLParam(r, "hotelSlug"),
	}

	if errMsg := validation.CheckErrors(pathParams, validation.CustomValidationError); errMsg != nil {
		return models.HotelRef{}, errMsg
	}

	return mapper.HotelPathParamsToEntity(pathParams), nil
}

func ParseUUIDParam(r *http.Request, paramName string) (uuid.UUID, error) {
	paramID := chi.URLParam(r, paramName)
	id, err := uuid.Parse(paramID)
	if err != nil {
		return uuid.Nil, consts.ErrInvalidHotelID
	}
	return id, nil
}
