package service

import (
	"context"

	"hotel/internal/repository/models"

	"github.com/google/uuid"
)

func (s *Service) RoomCreate(ctx context.Context, hotel models.HotelRef, room models.RoomCreate) (models.Room, error) {
	newRoom, err := s.repo.RoomCreate(ctx, hotel, room)
	if err != nil {
		return models.Room{}, err
	}

	return newRoom, nil
}

func (s *Service) RoomGetByID(ctx context.Context, hotel models.HotelRef, id uuid.UUID) (models.Room, error) {
	room, err := s.repo.RoomGetByID(ctx, hotel, id)
	if err != nil {
		return models.Room{}, err
	}

	return room, nil
}

func (s *Service) RoomGetAll(ctx context.Context, hotel models.HotelRef, page, limit uint64) (models.RoomList, error) {
	offset := (page - 1) * limit
	roomList, err := s.repo.RoomGetAll(ctx, hotel, limit, offset)
	if err != nil {
		return models.RoomList{}, err
	}

	return roomList, nil
}

func (s *Service) RoomUpdateByID(
	ctx context.Context,
	hotel models.HotelRef,
	id uuid.UUID,
	room models.RoomUpdate,
) error {
	if err := s.repo.RoomUpdateByID(ctx, hotel, id, room); err != nil {
		return err
	}

	return nil
}

func (s *Service) RoomStatusUpdateByID(
	ctx context.Context,
	hotel models.HotelRef,
	id uuid.UUID,
	room models.RoomStatusUpdate,
) error {
	if err := s.repo.RoomStatusUpdateByID(ctx, hotel, id, room); err != nil {
		return err
	}

	return nil
}

func (s *Service) RoomDeleteByID(ctx context.Context, hotel models.HotelRef, id uuid.UUID) error {
	if err := s.repo.RoomDeleteByID(ctx, hotel, id); err != nil {
		return err
	}

	return nil
}
