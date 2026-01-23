package helper

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"hotel/pkg/lib/utils/consts"
)

var (
	errBookingNotFound = status.Error(codes.NotFound, consts.HotelNotFound.Error())
	errInternalServer  = status.Error(codes.Internal, consts.ErrInternalServer.Error())
)

func DomainError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, consts.HotelNotFound):
		return errBookingNotFound

	default:
		return errInternalServer
	}
}
