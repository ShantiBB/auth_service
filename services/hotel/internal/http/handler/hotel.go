package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"fukuro-reserve/pkg/utils/consts"
	"fukuro-reserve/pkg/utils/helper"
	"hotel/internal/http/dto/request"
	"hotel/internal/http/dto/response"
	"hotel/internal/repository/postgres/models"
)

type HotelService interface {
	HotelCreate(ctx context.Context, h models.HotelCreate) (models.Hotel, error)
	HotelGetByIDOrName(ctx context.Context, field any) (models.Hotel, error)
	HotelGetAll(ctx context.Context, limit, offset uint64) (models.HotelList, error)
	HotelUpdateByID(ctx context.Context, id uuid.UUID, h models.HotelUpdate) error
	HotelDeleteByID(ctx context.Context, id uuid.UUID) error
}

// HotelCreate   godoc
// @Summary      Create hotel
// @Description  Create a new hotel from admin provider
// @Tags         hotels
// @Accept       json
// @Produce      json
// @Param        request  body      request.HotelCreate  true  "Hotel data"
// @Success      201      {object}  response.Hotel
// @Failure      400      {object}  response.ErrorSchema
// @Failure      401      {object}  response.ErrorSchema
// @Failure      409      {object}  response.ErrorSchema
// @Failure      500      {object}  response.ErrorSchema
// @Security     Bearer
// @Router       /hotels/  [post]
func (h *Handler) HotelCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.HotelCreate

	if err := helper.ParseJSON(w, r, &req, nil); err != nil {
		return
	}

	newHotel := h.HotelCreateRequestToEntity(req)
	createdHotel, err := h.svc.HotelCreate(ctx, newHotel)
	if err != nil {
		if errors.Is(err, consts.UniqueHotelField) {
			errMsg := response.ErrorResp(consts.UniqueHotelField)
			helper.SendError(w, r, http.StatusConflict, errMsg)
			return
		}
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	hotelResponse := h.HotelEntityToResponse(createdHotel)
	helper.SendSuccess(w, r, http.StatusCreated, hotelResponse)
}

// HotelGetAll    godoc
//
//	@Summary		Get hotels
//	@Description	Get hotels from admin or moderator provider
//	@Tags			hotels
//	@Accept			json
//	@Produce		json
//	@Param			page	query		uint64	false	"Page"	default(1)
//	@Param			limit	query		uint64	false	"Limit"	default(100)
//	@Success		200		{object}	response.HotelList
//	@Failure		401		{object}	response.ErrorSchema
//	@Failure		500		{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/hotels/ [get]
func (h *Handler) HotelGetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pagination, err := helper.ParsePaginationQuery(r)
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidQueryParam)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	hotelList, err := h.svc.HotelGetAll(ctx, pagination.Page, pagination.Limit)
	if err != nil {
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	hotels := make([]response.HotelShort, 0, len(hotelList.Hotels))
	for _, hotel := range hotelList.Hotels {
		hotelResponse := h.HotelShortEntityToShortResponse(hotel)
		hotels = append(hotels, hotelResponse)
	}

	totalPageCount := (hotelList.TotalCount + pagination.Limit - 1) / pagination.Limit
	pageLinks := helper.BuildPaginationLinks(r, pagination, totalPageCount)
	hotelListResp := response.HotelList{
		Hotels:           hotels,
		CurrentPage:      pagination.Page,
		Limit:            pagination.Limit,
		Links:            pageLinks,
		TotalPageCount:   totalPageCount,
		TotalHotelsCount: hotelList.TotalCount,
	}

	helper.SendSuccess(w, r, http.StatusOK, hotelListResp)
}

// HotelGetByID    godoc
//
//	@Summary		Get hotel by ID
//	@Description	Get hotel by ID from admin, moderator or owner provider
//	@Tags			hotels
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Hotel ID"
//	@Success		200	{object}	response.Hotel
//	@Failure		400	{object}	response.ErrorSchema
//	@Failure		401	{object}	response.ErrorSchema
//	@Failure		404	{object}	response.ErrorSchema
//	@Failure		500	{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/hotels/{id} [get]
func (h *Handler) HotelGetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paramID := chi.URLParam(r, "id")
	id, err := uuid.Parse(paramID)
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidID)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	hotel, err := h.svc.HotelGetByIDOrName(ctx, id)
	if err != nil {
		if errors.Is(err, consts.HotelNotFound) {
			errMsg := response.ErrorResp(consts.HotelNotFound)
			helper.SendError(w, r, http.StatusNotFound, errMsg)
			return
		}
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	hotelResponse := h.HotelEntityToResponse(hotel)
	helper.SendSuccess(w, r, http.StatusOK, hotelResponse)
}

// HotelUpdateByID    godoc
//
//	@Summary		Update hotel by ID
//	@Description	Update hotel by ID from admin, moderator or owner provider
//	@Tags			hotels
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Hotel ID"
//
// @Param           request  body   request.HotelUpdate  true  "Hotel data"
//
//	@Success		200	{object}	response.HotelUpdate
//	@Failure		400	{object}	response.ErrorSchema
//	@Failure		401	{object}	response.ErrorSchema
//	@Failure		404	{object}	response.ErrorSchema
//	@Failure		500	{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/hotels/{id} [put]
func (h *Handler) HotelUpdateByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paramID := chi.URLParam(r, "id")
	id, err := uuid.Parse(paramID)
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidID)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	var req request.HotelUpdate
	if err = helper.ParseJSON(w, r, &req, nil); err != nil {
		return
	}

	hotelUpdate := h.HotelUpdateRequestToEntity(req)
	if err = h.svc.HotelUpdateByID(ctx, id, hotelUpdate); err != nil {
		if errors.Is(err, consts.HotelNotFound) {
			errMsg := response.ErrorResp(consts.HotelNotFound)
			helper.SendError(w, r, http.StatusNotFound, errMsg)
			return
		}
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	hotelResponse := h.HotelUpdateEntityToResponse(id, hotelUpdate)
	helper.SendSuccess(w, r, http.StatusOK, hotelResponse)
}

// HotelDeleteByID    godoc
//
//	@Summary		Delete hotel by ID
//	@Description	Delete hotel by ID from admin or owner provider
//	@Tags			hotels
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Hotel ID"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	response.ErrorSchema
//	@Failure		401	{object}	response.ErrorSchema
//	@Failure		404	{object}	response.ErrorSchema
//	@Failure		500	{object}	response.ErrorSchema
//	@Security		Bearer
//	@Router			/hotels/{id} [delete]
func (h *Handler) HotelDeleteByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paramID := chi.URLParam(r, "id")
	id, err := uuid.Parse(paramID)
	if err != nil {
		errMsg := response.ErrorResp(consts.InvalidID)
		helper.SendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	if err = h.svc.HotelDeleteByID(ctx, id); err != nil {
		if errors.Is(err, consts.HotelNotFound) {
			errMsg := response.ErrorResp(consts.HotelNotFound)
			helper.SendError(w, r, http.StatusNotFound, errMsg)
			return
		}
		errMsg := response.ErrorResp(consts.InternalServer)
		helper.SendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	helper.SendSuccess(w, r, http.StatusNoContent, nil)
}
