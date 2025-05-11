package exceptions

import "errors"

var (
	ErrNicknameNotFound = errors.New("nickname-not-found")
	ErrTokenIsEmpty     = errors.New("token-is-empty")
	ErrPlayerNotFound   = errors.New("player-not-found")
	ErrPlayerIDIsEmpty  = errors.New("player-id-is-empty")
)
