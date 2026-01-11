package service

type Repository interface {
	BookingTransactionRepository
	BookingRepository
	BookingRoomRepository
	RoomLockRepository
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo}
}
