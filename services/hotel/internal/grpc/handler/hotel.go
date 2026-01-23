package handler

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	hotelv1 "hotel/api/hotel/v1"
	"hotel/internal/grpc/utils/helper"
	"hotel/internal/grpc/utils/mapper"
)

func (h *Handler) CreateHotel(
	ctx context.Context,
	req *hotelv1.CreateHotelRequest,
) (*hotelv1.CreateHotelResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	hotel := mapper.CreateHotelRequestToDomain(req)

	created, err := h.svc.CreateHotel(ctx, hotel)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.DomainError(err)
	}

	return &hotelv1.CreateHotelResponse{
		Hotel: mapper.HotelResponseToProto(created),
	}, nil
}

func (h *Handler) GetHotels(
	ctx context.Context,
	req *hotelv1.GetHotelsRequest,
) (*hotelv1.GetHotelsResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	page, limit, hotelInfo := mapper.GetHotelsRequestToDomain(req)

	bookingList, err := h.svc.GetHotels(ctx, hotelInfo, "title", page, limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.DomainError(err)
	}

	return &hotelv1.GetHotelsResponse{
		Hotels:     mapper.HotelsResponseToProto(bookingList.Hotels),
		TotalCount: bookingList.TotalCount,
		Page:       req.Page,
		Limit:      req.Limit,
	}, nil
}
