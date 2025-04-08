package exceptions

import "errors"

var (
	ErrInvalidGameStatus = errors.New("invalid-game-status")
	ErrConvertJSON       = errors.New("invalid-request")
	ErrInvalidType       = errors.New("invalid-action-type")
	ErrInvalidRequest    = errors.New("invalid-request")
)
