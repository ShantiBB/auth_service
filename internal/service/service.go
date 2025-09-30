package service

import "auth_service/internal/database/postgres"

type Service struct {
	repo *postgres.Repository
}

func New(repo *postgres.Repository) *Service {
	return &Service{repo: repo}
}
