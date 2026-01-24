package consts

import "errors"

const (
	MsgUserNotFound        = "user not found"
	MsgInvalidEmail        = "invalid email format"
	MsgUniqueUserField     = "username or email already exists"
	MsgInvalidRole         = "invalid role status"
	MsgPasswordHashing     = "error hashing password"
	MsgInvalidPassword     = "minimum length 8 characters"
	MsgInvalidCredentials  = "invalid credentials"
	MsgInvalidRefreshToken = "invalid token"
	MsgUnauthorized        = "unauthorized"
	MsgFieldRequired       = "field is required"
	MsgInvalidID           = "invalid ID"
	MsgInvalidQueryParam   = "invalid query parameter"
	MsgInternalServer      = "internal server error"
	MsgForbidden           = "forbidden"
	MsgInvalidJSON         = "invalid JSON body"
)

var (
	ErrUserNotFound        = errors.New(MsgUserNotFound)
	ErrInvalidEmail        = errors.New(MsgInvalidEmail)
	ErrUniqueUserField     = errors.New(MsgUniqueUserField)
	ErrInvalidRole         = errors.New(MsgInvalidRole)
	ErrPasswordHashing     = errors.New(MsgPasswordHashing)
	ErrInvalidPassword     = errors.New(MsgInvalidPassword)
	ErrInvalidCredentials  = errors.New(MsgInvalidCredentials)
	ErrInvalidRefreshToken = errors.New(MsgInvalidRefreshToken)
	ErrUnauthorized        = errors.New(MsgUnauthorized)
	ErrFieldRequired       = errors.New(MsgFieldRequired)
	ErrInvalidID           = errors.New(MsgInvalidID)
	ErrInvalidQueryParam   = errors.New(MsgInvalidQueryParam)
	ErrInternalServer      = errors.New(MsgInternalServer)
	ErrForbidden           = errors.New(MsgForbidden)
	ErrInvalidJSON         = errors.New(MsgInvalidJSON)
)
