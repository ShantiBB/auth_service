package helper

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"booking/pkg/utils/consts"
)

func DomainError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, consts.BookingNotFound):
		return status.Error(codes.NotFound, consts.BookingNotFound.Error())

	case errors.Is(err, consts.BookingRoomNotFound):
		return status.Error(codes.NotFound, consts.BookingRoomNotFound.Error())

	case errors.Is(err, consts.RoomLockAlreadyExist):
		return status.Error(codes.AlreadyExists, consts.RoomLockAlreadyExist.Error())

	case errors.Is(err, consts.ErrPriceChanged):
		return status.Error(codes.FailedPrecondition, consts.ErrPriceChanged.Error())

	default:
		return status.Error(codes.Internal, consts.InternalServer.Error())
	}
}
