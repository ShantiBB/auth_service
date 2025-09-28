package service

import (
	"context"

	"auth_service/internal/entity"
)

type UserService interface {
	Create(ctx context.Context, user entity.UserCreate) (*entity.User, error)
	Get(ctx context.Context, id int64) (*entity.User, error)
	GetAll(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, user entity.User) (*entity.User, error)
	Delete(ctx context.Context, id int64) error
}

func (s *Service) Create(ctx context.Context, user entity.UserCreate) (*entity.User, error) {
	newUser, err := s.UserRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *Service) Get(ctx context.Context, id int64) (*entity.User, error) {
	user, err := s.UserRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetAll(ctx context.Context) ([]entity.User, error) {
	users, err := s.UserRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) Update(ctx context.Context, user entity.User) (*entity.User, error) {
	if err := s.UserRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	updatedUser, err := s.UserRepo.Get(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	err := s.UserRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
