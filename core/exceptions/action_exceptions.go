package exceptions

import "errors"

var (
	ErrInvalidAction  = errors.New("invalid-action")
	ErrActionNotFound = errors.New("action-not-found")
	ErrEmptyActionID  = errors.New("empty-action-id")
)
