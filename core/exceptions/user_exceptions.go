package exceptions

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrUserNotFound = errors.New("user-not-found")
	ErrInvalidUUID  = errors.New("invalid-uuid-format")
)
