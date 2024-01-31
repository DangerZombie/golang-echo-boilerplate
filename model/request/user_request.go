package request

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserProfileRequest struct {
	Id string `json:"id"`
}
