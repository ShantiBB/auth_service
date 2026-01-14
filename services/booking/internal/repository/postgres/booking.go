package postgres

import (
	"context"
	"errors"

	"booking/internal/repository/models"
	"booking/internal/repository/postgres/query"
	"booking/pkg/utils/consts"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateBooking(ctx context.Context, tx pgx.Tx, b models.CreateBooking) (models.Booking, error) {
	db := r.executor(tx)

	newBooking := b.ToRead()
	insertArgs := []any{
		b.UserID,
		b.HotelID,
		b.CheckIn,
		b.CheckOut,
		b.GuestName,
		b.GuestEmail,
		b.GuestPhone,
		b.Currency,
		b.ExpectedTotalAmount,
		b.FinalTotalAmount,
	}
	scanArgs := []any{
		&newBooking.ID,
		&newBooking.Status,
		&newBooking.CreatedAt,
		&newBooking.UpdatedAt,
	}

	if err := db.QueryRow(ctx, query.CreateBooking, insertArgs...).Scan(scanArgs...); err != nil {
		return models.Booking{}, err
	}

	return newBooking, nil
}

func (r *Repository) GetBookingsByHotelInfo(
	ctx context.Context,
	tx pgx.Tx,
	bookingRef models.BookingRef,
	limit uint64,
	offset uint64,
) (models.BookingList, error) {
	db := r.executor(tx)

	var bookingList models.BookingList
	selectArgs := []any{
		bookingRef.UserID,
		bookingRef.HotelID,
		bookingRef.Status,
		limit,
		offset,
	}

	rows, err := db.Query(ctx, query.GetBookingsByHotelInfo, selectArgs...)
	if err != nil {
		return models.BookingList{}, err
	}

	var b models.BookingShort
	for rows.Next() {
		err = rows.Scan(
			&b.ID,
			&b.UserID,
			&b.HotelID,
			&b.CheckIn,
			&b.CheckOut,
			&b.Status,
			&b.GuestName,
			&b.GuestEmail,
			&b.GuestPhone,
			&b.Currency,
			&b.ExpectedTotalAmount,
			&b.FinalTotalAmount,
		)
		if err != nil {
			return models.BookingList{}, err
		}

		bookingList.Bookings = append(bookingList.Bookings, b)
	}

	if err = db.
		QueryRow(
			ctx,
			query.GetBookingCountRows,
			bookingRef.UserID,
			bookingRef.HotelID,
			bookingRef.Status,
		).
		Scan(&bookingList.TotalCount); err != nil {
		return models.BookingList{}, err
	}

	return bookingList, nil
}

func (r *Repository) GetBookingByID(ctx context.Context, tx pgx.Tx, bookingID uuid.UUID) (models.Booking, error) {
	db := r.executor(tx)

	var b models.Booking
	scanArgs := []any{
		&b.ID,
		&b.UserID,
		&b.HotelID,
		&b.CheckIn,
		&b.CheckOut,
		&b.Status,
		&b.GuestName,
		&b.GuestEmail,
		&b.GuestPhone,
		&b.Currency,
		&b.ExpectedTotalAmount,
		&b.FinalTotalAmount,
		&b.CreatedAt,
		&b.UpdatedAt,
	}

	if err := db.QueryRow(ctx, query.GetBookingByID, bookingID).Scan(scanArgs...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Booking{}, consts.BookingNotFound
		}
		return models.Booking{}, err
	}

	return b, nil
}

func (r *Repository) UpdateBookingGuestInfoByID(
	ctx context.Context, tx pgx.Tx, id uuid.UUID, b models.UpdateBooking,
) error {
	db := r.executor(tx)

	updateArgs := []any{
		id,
		b.GuestName,
		b.GuestEmail,
		b.GuestPhone,
	}

	row, err := db.Exec(ctx, query.UpdateBookingGuestInfoByID, updateArgs...)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingNotFound
	}

	return nil
}

func (r *Repository) UpdateBookingStatusByID(
	ctx context.Context, tx pgx.Tx, id uuid.UUID, b models.BookingStatusInfo,
) error {
	db := r.executor(tx)

	row, err := db.Exec(ctx, query.UpdateBookingStatusByID, id, b.Status)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingNotFound
	}

	return nil
}

func (r *Repository) DeleteBookingByID(ctx context.Context, tx pgx.Tx, id uuid.UUID) error {
	db := r.executor(tx)

	row, err := db.Exec(ctx, query.DeleteBookingByID, id)
	if err != nil {
		return err
	}
	if rowAffected := row.RowsAffected(); rowAffected == 0 {
		return consts.BookingNotFound
	}

	return nil
}
