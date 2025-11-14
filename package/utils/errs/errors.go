package errs

import "errors"

var (
	InternalServer     = errors.New("internal server error")
	UserNotFound       = errors.New("user not found")
	UserRetrieving     = errors.New("error retrieving users")
	UniqueUserField    = errors.New("username or email already exists")
	Unauthorized       = errors.New("unauthorized")
	Forbidden          = errors.New("forbidden")
	InvalidCredentials = errors.New("invalid credentials")
	InvalidID          = errors.New("invalid user ID")
	InvalidToken       = errors.New("invalid token")
	InvalidJSON        = errors.New("invalid JSON body")
	PasswordHashing    = errors.New("error hashing password")
)
