package request

// swagger:parameters InsertDriverRequest
type InsertDriverRequest struct {
	// Name
	// in: string
	// example: John
	Name string `json:"name"`

	// License Number
	// in: string
	// example: 8901119012
	LicenseNumber string `json:"license_number"`

	// Is Available
	// in: bool
	// example: true
	IsAvailable bool `json:"is_available"`
}

type GetListDriversRequest struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Sort  string `json:"sort"`
	Dir   string `json:"dir"`
	Name  string `json:"name"`
}

type GetDriverByNumber struct {
	Number string `json:"number"`
}

type UpdateDriverByNumber struct {
	Number      string `json:"number"`
	IsAvailable bool   `json:"is_available,omitempty"`
}

type DeleteDriverByNumber struct {
	Number string `json:"number"`
}
