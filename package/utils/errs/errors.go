package errs

import "errors"

var (
	UserNotFound       = errors.New("user not found")
	UniqueUserField    = errors.New("username or email already exists")
	InvalidCredentials = errors.New("invalid credentials")
	InvalidToken       = errors.New("invalid token")
)
