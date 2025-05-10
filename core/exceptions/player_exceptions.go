package exceptions

import "errors"

var (
	ErrNicknameNotFound = errors.New("nickname-not-found")
	ErrTokenIsEmpty     = errors.New("token-is-empty")
)
