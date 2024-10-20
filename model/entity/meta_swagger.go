package entity

// swagger:model MetaSingleSuccessResponse
type MetaSingleSuccessResponse struct {
	// Code of Response
	// in: int
	// example: 1000
	Code int `json:"code"`

	// Message of Response
	// in: string
	// example: Success
	Message string `json:"message"`
}

// swagger:response
type MetaSingleResponse struct {
	// in: body
	Response MetaSingleSuccessResponse `json:"response"`
}
