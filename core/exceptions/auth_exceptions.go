package exceptions

import "errors"

var (
	ErrNoState         = errors.New("no-state-founded")
	ErrInvalidState    = errors.New("state-are-not-match")
	ErrCodeNotFound    = errors.New("code-not-founded")
	ErrExchangeFailed  = errors.New("code-token-exchange-failed")
	ErrFetchGoogleUser = errors.New("fetch-user-failed")
)
