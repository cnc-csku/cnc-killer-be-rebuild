package responses

type UserResponse struct {
	IdToken string `json:"idToken"`
	Role    string `json:"role"`
}

type RoleResponse struct {
	Role string `json:"role"`
}
