package service

import (
	"context"

	"github.com/google/uuid"

	"hotel/internal/repository/postgres/models"
)

type HotelRepository interface {
	HotelCreate(ctx context.Context, h models.HotelCreate) (models.Hotel, error)
	HotelGetByIDOrName(ctx context.Context, field any) (models.Hotel, error)
	HotelGetAll(ctx context.Context, limit, offset uint64) (models.HotelList, error)
	HotelUpdateByID(ctx context.Context, id uuid.UUID, h models.HotelUpdate) error
	HotelDeleteByID(ctx context.Context, id uuid.UUID) error
}

func (s *Service) HotelCreate(ctx context.Context, h models.HotelCreate) (models.Hotel, error) {
	newHotel, err := s.repo.HotelCreate(ctx, h)
	if err != nil {
		return models.Hotel{}, err
	}

	return newHotel, nil
}

func (s *Service) HotelGetByIDOrName(ctx context.Context, field any) (models.Hotel, error) {
	h, err := s.repo.HotelGetByIDOrName(ctx, field)
	if err != nil {
		return models.Hotel{}, err
	}

	return h, nil
}

func (s *Service) HotelGetAll(ctx context.Context, page, limit uint64) (models.HotelList, error) {
	offset := (page - 1) * limit
	hotelList, err := s.repo.HotelGetAll(ctx, limit, offset)
	if err != nil {
		return models.HotelList{}, err
	}

	return hotelList, nil
}

func (s *Service) HotelUpdateByID(ctx context.Context, id uuid.UUID, h models.HotelUpdate) error {
	if err := s.repo.HotelUpdateByID(ctx, id, h); err != nil {
		return err
	}

	return nil
}

func (s *Service) HotelDeleteByID(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.HotelDeleteByID(ctx, id); err != nil {
		return err
	}

	return nil
}
