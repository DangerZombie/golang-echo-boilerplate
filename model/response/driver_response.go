package response

// swagger:model InsertDriverResponse
type InsertDriverResponse struct {
	// ID
	// in: string
	// example: 5c091047-f19c-4235-be8d-b6596aecf880
	Id            string `json:"id"`

	// Name
	// in: string
	// example: John
	Name          string `json:"name"`

	// Licensse Number
	// in: string
	// example: 89011231091
	LicenseNumber string `json:"license_number"`

	// Is Available
	// in: bool
	// example: true
	IsAvailable   bool   `json:"is_available"`

	// Created At
	// in: string
	// example: 2023-03-30 10:39:44
	CreatedAt     string `json:"created_at"`

	// Created By
	// in: string
	// example: system
	CreatedBy     string `json:"created_by"`

	// Updated At
	// in: string
	// example: 2023-03-30 10:39:44
	UpdatedAt     string `json:"updated_at"`

	// Updated By
	// in: string
	// example: system
	UpdatedBy     string `json:"updated_by"`
}

type GetListDriversResponse struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	LicenseNumber string `json:"license_number"`
	IsAvailable   bool   `json:"is_available"`
	CreatedAt     string `json:"created_at"`
	CreatedBy     string `json:"created_by"`
	UpdatedAt     string `json:"updated_at"`
	UpdatedBy     string `json:"updated_by"`
}

type GetDriverByNumberResponse struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	LicenseNumber string `json:"license_number"`
	IsAvailable   bool   `json:"is_available"`
	CreatedAt     string `json:"created_at"`
	CreatedBy     string `json:"created_by"`
	UpdatedAt     string `json:"updated_at"`
	UpdatedBy     string `json:"updated_by"`
}

type UpdateDriverByNumberResponse struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	LicenseNumber string `json:"license_number"`
	IsAvailable   bool   `json:"is_available"`
	CreatedAt     string `json:"created_at"`
	CreatedBy     string `json:"created_by"`
	UpdatedAt     string `json:"updated_at"`
	UpdatedBy     string `json:"updated_by"`
}

type DeleteDriverByNumberResponse struct {
	Message string `json:"message"`
}
