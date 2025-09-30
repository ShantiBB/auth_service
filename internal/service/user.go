package service

import (
	"context"

	"auth_service/internal/domain/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.UserCreate) (*models.User, error)
	GetUser(ctx context.Context, id int64) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id int64) error
}

func (s *Service) CreateUser(ctx context.Context, user models.UserCreate) (*models.User, error) {
	newUser, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *Service) GetUser(ctx context.Context, id int64) (*models.User, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetAllUsers(ctx context.Context) ([]models.User, error) {
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	if err := s.repo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	updatedUser, err := s.repo.GetUser(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *Service) DeleteUser(ctx context.Context, id int64) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
