package models

type User struct {
	Email        string  `db:"email"`
	Role         string  `db:"user_role"`
	RefreshToken *string `db:"refresh_token"`
}

var RoleEnum = map[string]bool{
	"admin": true,
	"user":  true,
}

type roleValue struct {
	Admin string
	User  string
}

var UserRoles = &roleValue{
	Admin: "admin",
	User:  "user",
}
