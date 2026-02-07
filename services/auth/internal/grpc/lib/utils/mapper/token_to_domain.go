package mapper

import (
	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
	"github.com/ShantiBB/fukuro-reserve/services/auth/pkg/lib/utils/jwt"
)

func RefreshTokenRequestToDomain(req *userv1.RefreshTokenRequest) *jwt.Token {
	return &jwt.Token{
		Refresh: req.RefreshToken,
	}
}
