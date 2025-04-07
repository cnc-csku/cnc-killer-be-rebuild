package models

type User struct {
	UserID  string `db:"user_id"`
	IdToken string `db:"id_token"`
	Role    string `db:"user_role"`
}

var RoleEnum = map[string]bool{
	"admin": true,
	"user":  true,
}
