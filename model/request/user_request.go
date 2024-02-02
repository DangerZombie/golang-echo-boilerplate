package request

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserProfileRequest struct {
	Id string `json:"id"`
}

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}
