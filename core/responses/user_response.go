package responses

type UserResponse struct {
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
}

type RoleResponse struct {
	Role string `json:"role"`
}
