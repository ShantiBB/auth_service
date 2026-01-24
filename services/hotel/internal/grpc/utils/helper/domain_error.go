package helper

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"hotel/pkg/lib/utils/consts"
)

var (
	errHotelNotFound    = status.Error(codes.NotFound, consts.MsgHotelNotFound)
	errRoomNotFound     = status.Error(codes.NotFound, consts.MsgRoomNotFound)
	errUniqueHotelField = status.Error(codes.NotFound, consts.MsgUniqueHotelField)
	errUniqueRoomField  = status.Error(codes.NotFound, consts.MsgUniqueRoomField)
	errInternalServer   = status.Error(codes.Internal, consts.MsgInternalServer)
)

func DomainError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, consts.ErrHotelNotFound):
		return errHotelNotFound
	case errors.Is(err, consts.ErrRoomNotFound):
		return errRoomNotFound
	case errors.Is(err, consts.ErrUniqueHotelField):
		return errUniqueHotelField
	case errors.Is(err, consts.ErrUniqueRoomField):
		return errUniqueRoomField

	default:
		return errInternalServer
	}
}
