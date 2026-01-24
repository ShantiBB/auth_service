package service

import (
	"context"
	"errors"
	"log/slog"

	"auth/internal/grpc/utils/helper"
	"auth/internal/repository/models"
	"auth/pkg/lib/utils/consts"
	jwt2 "auth/pkg/lib/utils/jwt"
)

func (s *Service) RegisterByEmail(ctx context.Context, user *models.CreateUser) (*jwt2.Token, error) {
	created, err := s.repo.InsertUser(ctx, user)
	if err != nil {
		slog.Error("failed create user", "err:", err.Error())
		return nil, err
	}

	return jwt2.GenerateAllTokens(created.ID, created.Role, s.tokenCreds)
}

func (s *Service) LoginByEmail(ctx context.Context, user *models.CreateUser) (*jwt2.Token, error) {
	userCred, err := s.repo.SelectUserCredentialsByEmail(ctx, user.Email)
	if err != nil {
		slog.Error("failed login user", "err:", err.Error())
		return nil, err
	}

	if !helper.VerifyPassword(user.Password, userCred.Password) {
		return nil, consts.ErrInvalidCredentials
	}

	return jwt2.GenerateAllTokens(userCred.ID, userCred.Role, s.tokenCreds)
}

func (s *Service) RefreshToken(token *jwt2.Token) (*jwt2.Token, error) {
	claims, err := jwt2.GetClaimsRefreshToken(s.tokenCreds.RefreshSecret, token.Refresh)
	if err != nil {
		if errors.Is(err, consts.ErrInvalidRefreshToken) {
			return nil, consts.ErrInvalidRefreshToken
		}
		return nil, err
	}

	access, err := jwt2.GenerateAccessToken(claims.Sub, claims.Role, s.tokenCreds)
	if err != nil {
		return nil, err
	}

	token.Access = access
	return token, nil
}
