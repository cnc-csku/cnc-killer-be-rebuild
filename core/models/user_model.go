package models

type User struct {
	Email        string  `db:"email"`
	Nickname     *string `db:"nickname"`
	Role         string  `db:"user_role"`
	RefreshToken *string `db:"refresh_token"`
}

type JWTToken struct {
	Email string `mapstructure:"email"`
	Role  string `mapstructure:"role"`
	Exp   int64  `mapstructure:"exp"`
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
