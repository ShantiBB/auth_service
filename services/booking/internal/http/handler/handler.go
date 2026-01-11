package handler

type Service interface {
	BookingService
}

type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc}
}
