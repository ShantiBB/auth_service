package handler

type Service interface {
	HotelService
	RoomService
}

type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc}
}
