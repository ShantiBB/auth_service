package mapper

import (
	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
	"github.com/ShantiBB/fukuro-reserve/services/auth/internal/grpc/lib/utils/helper"
	"github.com/ShantiBB/fukuro-reserve/services/auth/internal/repository/models"
	"github.com/ShantiBB/fukuro-reserve/services/auth/pkg/lib/utils/consts"
)

func userRoleToDomain(role userv1.UserRole) models.UserRole {
	var s models.UserRole
	switch role {
	case userv1.UserRole_USER_ROLE_USER:
		s = models.UserRoleUser
	case userv1.UserRole_USER_ROLE_MODERATOR:
		s = models.UserRoleModerator
	case userv1.UserRole_USER_ROLE_ADMIN:
		s = models.UserRoleAdmin
	default:
		s = models.UserRoleUnspecified
	}
	return s
}

func CreateUserRequestToDomain(req *userv1.CreateUserRequest) (*models.CreateUser, error) {
	hashPass, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, consts.ErrPasswordHashing
	}
	return &models.CreateUser{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPass,
	}, nil
}

func UpdateUserRequestToDomain(req *userv1.UpdateUserRequest) *models.UpdateUser {
	return &models.UpdateUser{
		ID:       req.Id,
		Username: req.Username,
		Email:    req.Email,
	}
}

func UpdateUserRoleRequestToDomain(req *userv1.UpdateUserRoleRequest) models.UserRole {
	return userRoleToDomain(req.Role)
}

func RegisterUserRequestToDomain(req *userv1.RegisterUserRequest) (*models.CreateUser, error) {
	hashPass, err := helper.HashPassword(req.GetPassword())
	if err != nil {
		return nil, consts.ErrPasswordHashing
	}
	return &models.CreateUser{
		Email:    req.Email,
		Password: hashPass,
	}, nil
}

func LoginUserRequestToDomain(req *userv1.LoginUserRequest) *models.CreateUser {
	return &models.CreateUser{
		Email:    req.Email,
		Password: req.Password,
	}
}
