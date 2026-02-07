package handler

import (
	"context"

	"buf.build/go/protovalidate"

	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
	"github.com/ShantiBB/fukuro-reserve/services/auth/internal/repository/models"
	"github.com/ShantiBB/fukuro-reserve/services/auth/pkg/lib/utils/jwt"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.CreateUser) (*models.User, error)
	GetUsers(ctx context.Context, page, limit uint64) (*models.UserList, error)
	GetUserByID(ctx context.Context, id int64) (*models.User, error)
	UpdateUserByID(ctx context.Context, user *models.UpdateUser) error
	UpdateUserRoleStatus(ctx context.Context, id int64, role models.UserRole) error
	UpdateUserActiveStatus(ctx context.Context, id int64, status bool) error
	DeleteUserByID(ctx context.Context, id int64) error
}

type TokenService interface {
	RegisterByEmail(ctx context.Context, user *models.CreateUser) (*jwt.Token, error)
	LoginByEmail(ctx context.Context, user *models.CreateUser) (*jwt.Token, error)
	RefreshToken(token *jwt.Token) (*jwt.Token, error)
}

type Service interface {
	UserService
	TokenService
}

type Handler struct {
	userv1.UnimplementedUserServiceServer
	userv1.UnimplementedTokenServiceServer
	svc       Service
	validator protovalidate.Validator
}

func New(svc Service, validator protovalidate.Validator) *Handler {
	return &Handler{svc: svc, validator: validator}
}
