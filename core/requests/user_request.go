package requests

type UserLoginRequest struct {
	Password string `json:"password"`
}

type UserGetRoleRequest struct {
	UserID string `json:"user_id"`
}

type ChangeNicknameRequest struct {
	Nickname string `json:"nickname"`
}
