package errs

import "errors"

var (
	InternalServer     = errors.New("internal server error")
	UserNotFound       = errors.New("user not found")
	UniqueUserField    = errors.New("username or email already exists")
	InvalidCredentials = errors.New("invalid credentials")
	InvalidToken       = errors.New("invalid token")
	PasswordHashing    = errors.New("error hashing password")
)
