package service

import "auth_service/internal/database/postgres"

type Service struct {
	UserRepo postgres.UserRepository
}

func New(userRepo postgres.UserRepository) *Service {
	return &Service{
		UserRepo: userRepo,
	}
}
