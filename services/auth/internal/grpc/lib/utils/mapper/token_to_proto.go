package mapper

import (
	userv1 "github.com/ShantiBB/fukuro-reserve/services/auth/api/user/v1"
	"github.com/ShantiBB/fukuro-reserve/services/auth/pkg/lib/utils/jwt"
)

func JWTTokenResponseToProto(resp *jwt.Token) *userv1.Tokens {
	return &userv1.Tokens{
		Access:  resp.Access,
		Refresh: resp.Refresh,
	}
}
