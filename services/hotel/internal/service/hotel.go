package service

import (
	"context"

	"hotel/internal/repository/models"

	"github.com/gosimple/slug"
)

type HotelRepository interface {
	HotelCreate(
		ctx context.Context,
		hotelRef models.HotelRef,
		h models.HotelCreate,
	) (models.Hotel, error)
	HotelGetAll(
		ctx context.Context,
		hotelRef models.HotelRef,
		sortField string,
		limit uint64,
		offset uint64,
	) (models.HotelList, error)
	HotelGetBySlug(ctx context.Context, hotelRef models.HotelRef) (models.Hotel, error)
	HotelUpdateBySlug(ctx context.Context, hotelRef models.HotelRef, h models.HotelUpdate) error
	HotelTitleUpdateBySlug(
		ctx context.Context,
		hotelRef models.HotelRef,
		h models.HotelTitleUpdate,
	) error
	HotelDeleteBySlug(ctx context.Context, hotelRef models.HotelRef) error
}

func (s *Service) HotelCreate(
	ctx context.Context,
	hotel models.HotelRef,
	h models.HotelCreate,
) (models.Hotel, error) {
	hotelRef := models.HotelRef{
		CountryCode: hotel.CountryCode,
		CitySlug:    hotel.CitySlug,
	}
	h.Slug = slug.Make(h.Title)

	newHotel, err := s.repo.HotelCreate(ctx, hotelRef, h)
	if err != nil {
		return models.Hotel{}, err
	}

	return newHotel, nil
}

func (s *Service) HotelGetAll(
	ctx context.Context,
	hotel models.HotelRef,
	sortField string,
	page uint64,
	limit uint64,
) (models.HotelList, error) {
	hotelRef := models.HotelRef{
		CountryCode: hotel.CountryCode,
		CitySlug:    hotel.CitySlug,
	}
	offset := (page - 1) * limit

	hotelList, err := s.repo.HotelGetAll(ctx, hotelRef, sortField, limit, offset)
	if err != nil {
		return models.HotelList{}, err
	}

	return hotelList, nil
}

func (s *Service) HotelGetBySlug(ctx context.Context, hotel models.HotelRef) (models.Hotel, error) {
	hotelRef := models.HotelRef{
		CountryCode: hotel.CountryCode,
		CitySlug:    hotel.CitySlug,
		HotelSlug:   hotel.HotelSlug,
	}

	h, err := s.repo.HotelGetBySlug(ctx, hotelRef)
	if err != nil {
		return models.Hotel{}, err
	}

	return h, nil
}

func (s *Service) HotelUpdateBySlug(ctx context.Context, hotel models.HotelRef, h models.HotelUpdate) error {
	hotelRef := models.HotelRef{
		CountryCode: hotel.CountryCode,
		CitySlug:    hotel.CitySlug,
		HotelSlug:   hotel.HotelSlug,
	}

	if err := s.repo.HotelUpdateBySlug(ctx, hotelRef, h); err != nil {
		return err
	}

	return nil
}

func (s *Service) HotelTitleUpdateBySlug(
	ctx context.Context,
	hotel models.HotelRef,
	h models.HotelTitleUpdate,
) (models.HotelTitleUpdate, error) {
	hotelRef := models.HotelRef{
		CountryCode: hotel.CountryCode,
		CitySlug:    hotel.CitySlug,
		HotelSlug:   hotel.HotelSlug,
	}
	h.Slug = slug.Make(h.Title)

	if err := s.repo.HotelTitleUpdateBySlug(ctx, hotelRef, h); err != nil {
		return models.HotelTitleUpdate{}, err
	}

	return h, nil
}

func (s *Service) HotelDeleteBySlug(ctx context.Context, hotel models.HotelRef) error {
	hotelRef := models.HotelRef{
		CountryCode: hotel.CountryCode,
		CitySlug:    hotel.CitySlug,
		HotelSlug:   hotel.HotelSlug,
	}

	if err := s.repo.HotelDeleteBySlug(ctx, hotelRef); err != nil {
		return err
	}

	return nil
}
