package errs

import "errors"

var (
	UserCreateError    = errors.New("error creating user")
	UserNotFound       = errors.New("user not found")
	UniqueUserField    = errors.New("username or email already exists")
	InvalidCredentials = errors.New("invalid credentials")
	InvalidToken       = errors.New("invalid token")
	PasswordHashing    = errors.New("error hashing password")
)
