package response

// swagger:model LoginResponse
type LoginResponse struct {
	// Token
	// in: string
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MDY1ODMwNTYsInVzZXIiOiJhZG1pbiJ9.F-bBdILVQIg9kj8mWGn5ma7qDoyzSbiUojQz6EW_hJs
	Token string `json:"token"`
}

// swagger:model UserProfileResponse
type UserProfileResponse struct {
	// Id of user
	// in: string
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	Id string `json:"id,omitempty"`

	// Nickname of user
	// in: string
	// example: John Doe
	Nickname string `json:"Nickname,omitempty"`
}

// swagger:model RegisterUserResponse
type RegisterUserResponse struct {
	// Id of user
	// in: string
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	Id string `json:"id"`
}
