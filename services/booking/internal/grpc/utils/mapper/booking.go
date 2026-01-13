package mapper

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/timestamppb"

	bookingv1 "booking/api/booking/v1"
	"booking/internal/repository/models"
	"booking/pkg/utils/consts"
)

func CreateBookingRequestToDomain(req *bookingv1.CreateBookingRequest) (*models.CreateBooking, error) {
	hotelID, err := uuid.Parse(req.HotelId)
	if err != nil {
		return nil, consts.InvalidHotelID
	}

	expectedTotalAmount, err := decimal.NewFromString(req.ExpectedTotalAmount)
	if err != nil {
		return nil, consts.InvalidExpectedTotalAmountID
	}

	b := &models.CreateBooking{
		UserID:              req.UserId,
		HotelID:             hotelID,
		CheckIn:             req.CheckIn.AsTime(),
		CheckOut:            req.CheckOut.AsTime(),
		GuestName:           req.GuestName,
		GuestEmail:          req.GuestEmail,
		GuestPhone:          req.GuestPhone,
		Currency:            req.Currency,
		ExpectedTotalAmount: expectedTotalAmount,
	}

	return b, nil
}

func BookingToProto(b *models.Booking) *bookingv1.Booking {
	p := &bookingv1.Booking{
		Id:               b.ID.String(),
		UserId:           b.UserID,
		HotelId:          b.HotelID.String(),
		CheckIn:          timestamppb.New(b.CheckIn),
		CheckOut:         timestamppb.New(b.CheckOut),
		Status:           BookingStatusToProto(b.Status),
		GuestName:        b.GuestName,
		GuestEmail:       b.GuestEmail,
		GuestPhone:       b.GuestPhone,
		Currency:         b.Currency,
		FinalTotalAmount: b.FinalTotalAmount.String(),
		CreatedAt:        timestamppb.New(b.CreatedAt),
		UpdatedAt:        timestamppb.New(b.UpdatedAt),
	}

	if len(b.BookingRooms) > 0 {
		p.BookingRooms = make([]*bookingv1.BookingRoomResponse, 0, len(b.BookingRooms))
		for i := range b.BookingRooms {
			p.BookingRooms = append(p.BookingRooms, BookingRoomToProto(&b.BookingRooms[i]))
		}
	}

	return p
}

func GetBookingsRequestToDomain(req *bookingv1.GetBookingsRequest) (models.BookingRef, error) {
	bookingRef := models.BookingRef{
		UserID: req.UserId,
		Status: BookingStatusToDomain(req.Status),
	}

	hotelID, err := uuid.Parse(req.HotelId)
	if err != nil {
		return models.BookingRef{}, consts.InvalidHotelID
	}
	bookingRef.HotelID = hotelID

	return bookingRef, nil
}

func BookingListToProto(bookings []models.BookingShort) []*bookingv1.BookingShort {
	result := make([]*bookingv1.BookingShort, 0, len(bookings))
	for _, b := range bookings {
		result = append(result, BookingShortToProto(&b))
	}
	return result
}

func BookingShortToProto(b *models.BookingShort) *bookingv1.BookingShort {
	return &bookingv1.BookingShort{
		Id:                  b.ID.String(),
		UserId:              b.UserID,
		HotelId:             b.HotelID.String(),
		CheckIn:             timestamppb.New(b.CheckIn),
		CheckOut:            timestamppb.New(b.CheckOut),
		Status:              BookingStatusToProto(b.Status),
		GuestName:           b.GuestName,
		GuestEmail:          b.GuestEmail,
		GuestPhone:          b.GuestPhone,
		Currency:            b.Currency,
		ExpectedTotalAmount: b.ExpectedTotalAmount.String(),
		FinalTotalAmount:    b.FinalTotalAmount.String(),
		Rooms:               BookingRoomsShortToProto(b.BookingRooms),
	}
}

func BookingStatusToDomain(status bookingv1.BookingStatus) models.BookingStatus {
	var s models.BookingStatus
	switch status {
	case bookingv1.BookingStatus_BOOKING_STATUS_PENDING:
		s = models.BookingStatusPending
	case bookingv1.BookingStatus_BOOKING_STATUS_CONFIRMED:
		s = models.BookingStatusConfirmed
	case bookingv1.BookingStatus_BOOKING_STATUS_CANCELLED:
		s = models.BookingStatusCancelled
	default:
		return ""
	}
	return s
}

func BookingStatusToProto(s models.BookingStatus) bookingv1.BookingStatus {
	switch s {
	case models.BookingStatusPending:
		return bookingv1.BookingStatus_BOOKING_STATUS_PENDING
	case models.BookingStatusConfirmed:
		return bookingv1.BookingStatus_BOOKING_STATUS_CONFIRMED
	case models.BookingStatusCancelled:
		return bookingv1.BookingStatus_BOOKING_STATUS_CANCELLED
	default:
		return bookingv1.BookingStatus_BOOKING_STATUS_UNSPECIFIED
	}
}
