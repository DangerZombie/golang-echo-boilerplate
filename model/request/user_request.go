package request

// swagger:parameters LoginRequest
type LoginRequest struct {
	// in:body
	Body LoginRequestBody `json:"body"`
}

type LoginRequestBody struct {
	// Username of user
	// in: string
	// example: admin
	// required: true
	Username string `json:"username"`

	// Password of user
	// in: string
	// example: 123456
	// required: true
	Password string `json:"password"`
}

// swagger:parameters UserProfileRequest
type UserProfileRequest struct {
	// Id of user
	// in: string
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	Id string `json:"id"`
}

// swagger:parameters RegisterUserRequest
type RegisterUserRequest struct {
	// in:body
	Body RegisterUserRequestBody `json:"body"`
}

type RegisterUserRequestBody struct {
	// Username of user
	// in: string
	// example: admin
	// required: true
	Username string `json:"username"`

	// Password of user
	// in: string
	// example: 123456
	// required: true
	Password string `json:"password"`

	// Nickname of user will show as profile name
	// in: string
	// example: Administrator
	// required: true
	Nickname string `json:"nickname"`
}
