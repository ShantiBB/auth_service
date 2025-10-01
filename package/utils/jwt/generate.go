package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"auth_service/internal/config"
	"auth_service/internal/domain/entity"
)

func GenerateToken(sub int64, role string, ttl time.Duration, secret []byte) (string, error) {
	claims := entity.Claims{
		Sub:  sub,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func GenerateAccessToken(sub int64, role, AccessSecret string, AccessTTL time.Duration) (string, error) {
	return GenerateToken(sub, role, AccessTTL, []byte(AccessSecret))
}

func GenerateRefreshToken(sub int64, role, RefreshSecret string, RefreshTTL time.Duration) (string, error) {
	return GenerateToken(sub, role, RefreshTTL, []byte(RefreshSecret))
}

func GenerateAllTokens(sub int64, role string, cfg *config.Config) (*entity.Token, error) {
	var err error
	tokens := &entity.Token{}
	tokens.Access, err = GenerateAccessToken(sub, role, cfg.JWT.AccessSecret, cfg.JWT.AccessTokenTTL)
	if err != nil {
		return nil, err
	}

	tokens.Refresh, err = GenerateRefreshToken(sub, role, cfg.JWT.RefreshSecret, cfg.JWT.RefreshTokenTTL)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
