package repositories

type AuthRepository interface {
	GetAuthURL() error
	VerifyState(state string) bool
	GenerateState() (string, error)
}
