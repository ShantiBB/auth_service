package handler

import (
	"auth_service/internal/http/lib/schemas/request"
	"auth_service/package/utils/jwt"
)

var registerReq = request.Register{
	Email:    "test@example.com",
	Password: "password123",
}

var loginReq = request.LoginByEmail{
	Email:    "test@example.com",
	Password: "password123",
}

var refreshReq = request.RefreshToken{
	RefreshToken: "valid-refresh-token",
}

var tokensMock = jwt.Token{
	Access:  "access-token",
	Refresh: "refresh-token",
}
