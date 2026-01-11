package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"

	"booking/internal/repository/models"
	"booking/internal/service/utils/helper"
	"booking/pkg/utils/consts"

	"github.com/google/uuid"
)

type BookingTransactionRepository interface {
	BeginTx(ctx context.Context) (pgx.Tx, error)
}

type BookingRepository interface {
	CreateBooking(ctx context.Context, tx pgx.Tx, b models.CreateBooking) (models.Booking, error)
	GetBookingsByHotelInfo(
		ctx context.Context,
		tx pgx.Tx,
		bookingRef models.BookingRef,
		limit uint64,
		offset uint64,
	) (models.BookingList, error)
	GetBookingByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) (models.Booking, error)
	UpdateBookingGuestInfoByID(ctx context.Context, tx pgx.Tx, id uuid.UUID, b models.UpdateBooking) error
	UpdateBookingStatusByID(ctx context.Context, tx pgx.Tx, id uuid.UUID, b models.BookingStatusInfo) error
	DeleteBookingByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error
}

type BookingRoomRepository interface {
	CreateBookingRooms(ctx context.Context, tx pgx.Tx, rooms []models.CreateBookingRoom) ([]models.BookingRoom, error)
	GetBookingRoomsByBookingID(ctx context.Context, tx pgx.Tx, bookingID uuid.UUID) ([]models.BookingRoom, error)
	GetBookingRoomByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) (models.BookingRoom, error)
	UpdateBookingRoomGuestCountsByID(
		ctx context.Context,
		tx pgx.Tx,
		id uuid.UUID,
		bRoom models.BookingRoomGuestCounts,
	) error
	DeleteBookingRoomByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error
}

type RoomLockRepository interface {
	CreateRoomLocks(ctx context.Context, tx pgx.Tx, locks []models.CreateRoomLock) ([]models.RoomLock, error)
	GetRoomsLockByBookingID(ctx context.Context, tx pgx.Tx, bookingID uuid.UUID) ([]models.RoomLock, error)
	GetRoomLockByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) (models.RoomLock, error)
	UpdateRoomLockActivityByID(
		ctx context.Context,
		tx pgx.Tx,
		id uuid.UUID,
		roomLock models.UpdateRoomLockActivity,
	) error
	DeleteRoomLockByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error
}

func (s *Service) BookingCreate(
	ctx context.Context,
	b models.CreateBooking,
	rooms []models.CreateBookingRoom,
) (models.Booking, error) {
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return models.Booking{}, err
	}
	defer tx.Rollback(ctx)

	nights, err := helper.Nights(b.CheckIn, b.CheckOut)
	if err != nil {
		return models.Booking{}, err
	}

	finalTotalAmount := decimal.Zero
	for _, room := range rooms {
		roomTotal := room.PricePerNight.Mul(decimal.NewFromInt(int64(nights)))
		finalTotalAmount = finalTotalAmount.Add(roomTotal)
	}

	if !b.ExpectedTotalAmount.IsZero() {
		diff := finalTotalAmount.Sub(b.ExpectedTotalAmount).Abs()
		if diff.GreaterThan(decimal.RequireFromString("0.01")) {
			return models.Booking{}, consts.ErrPriceChanged
		}
	}

	b.FinalTotalAmount = finalTotalAmount

	newBooking, err := s.repo.CreateBooking(ctx, tx, b)
	if err != nil {
		return models.Booking{}, err
	}

	for i := range rooms {
		rooms[i].BookingID = newBooking.ID
	}

	newRooms, err := s.repo.CreateBookingRooms(ctx, tx, rooms)
	if err != nil {
		return models.Booking{}, err
	}

	now := time.Now()
	locks := make([]models.CreateRoomLock, 0, len(newRooms))
	for _, nr := range newRooms {
		locks = append(
			locks, models.CreateRoomLock{
				RoomID:    nr.RoomID,
				BookingID: newBooking.ID,
				StayRange: models.DateRange{
					Start: b.CheckIn,
					End:   b.CheckOut,
				},
				ExpiresAt: now.Add(15 * time.Minute),
			},
		)
	}

	if _, err = s.repo.CreateRoomLocks(ctx, tx, locks); err != nil {
		return models.Booking{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return models.Booking{}, err
	}

	return newBooking, nil
}

//func (s *Service) BookingGetAll(
//	ctx context.Context,
//	bookingRef models.BookingRef,
//	page uint64,
//	limit uint64,
//) (models.BookingList, error) {
//	offset := (page - 1) * limit
//
//	bookingList, err := s.repo.BookingGetAll(ctx, bookingRef, limit, offset)
//	if err != nil {
//		return models.BookingList{}, err
//	}
//
//	return bookingList, nil
//}
//
//func (s *Service) BookingGetByID(ctx context.Context, id uuid.UUID) (models.Booking, error) {
//	b, err := s.repo.BookingGetByID(ctx, id)
//	if err != nil {
//		return models.Booking{}, err
//	}
//
//	return b, nil
//}
//
//func (s *Service) BookingUpdateByID(ctx context.Context, id uuid.UUID, b models.UpdateBooking) error {
//	if err := s.repo.BookingUpdateByID(ctx, id, b); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *Service) BookingStatusUpdateByID(ctx context.Context, id uuid.UUID, b models.BookingStatusInfo) error {
//	if err := s.repo.BookingStatusUpdateByID(ctx, id, b); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *Service) BookingDeleteByID(ctx context.Context, id uuid.UUID) error {
//	if err := s.repo.BookingDeleteByID(ctx, id); err != nil {
//		return err
//	}
//
//	return nil
//}
