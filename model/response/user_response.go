package response

// swagger:model LoginResponse
type LoginResponse struct {
	// Token
	// in: string
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MDY1ODMwNTYsInVzZXIiOiJhZG1pbiJ9.F-bBdILVQIg9kj8mWGn5ma7qDoyzSbiUojQz6EW_hJs
	Token string `json:"token"`
}
