package exceptions

import "errors"

var (
	ErrUnauthorized  = errors.New("unauthorized")
	ErrUserNotFound  = errors.New("user-not-found")
	ErrEmailNotFound = errors.New("email-not-found")
	ErrInvalidUUID   = errors.New("invalid-uuid-format")
	ErrInvalidJWT    = errors.New("invalid-jwt")
)
