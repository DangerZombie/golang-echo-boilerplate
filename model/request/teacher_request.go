package request

// swagger:parameters TeacherCreateRequest
type TeacherCreateRequest struct {
	// in:body
	Body TeacherCreateRequestBody `json:"body"`
}

type TeacherCreateRequestBody struct {
	// User Id of the teacher
	// in: string
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	// required: true
	UserId string `json:"user_id"`

	// Job Title Id of the teacher
	// in: string
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	// required: true
	JobTitleId string `json:"job_title_id"`

	// Status of the teacher
	// in: string
	// example: [CONTRACT, ASSOCIATE, HONORER, PERMANENT]
	// required: true
	Status string `json:"status"`

	// Long time of experience in years
	// in: int
	// example: 12
	Experience int `json:"experience"`

	// Degree of the teacher
	// in: string
	// example: B.Ed
	Degree string `json:"degree"`

	Issuer string `json:"-"`
}

// swagger:parameters TeacherListRequest
type TeacherListRequest struct {
	// Page of the list
	// in: int
	// example: 1
	Page int `json:"page"`

	// Limit row of each page
	// in: int
	// example: 1
	Limit int `json:"limit"`

	// Sorting by column
	// in: string
	// example: name
	Sort string `json:"sort"`

	// Direction of sorting
	// in: string
	// example: asc
	Dir string `json:"dir"`

	// Filtering name
	// in: string
	// example: John Doe
	Name string `json:"name"`
}

// swagger:parameters TeacherDetailRequest
type TeacherDetailRequest struct {
	// Id of the teacher
	// in: path
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	Id string `json:"id"`
}

// swagger:parameters TeacherUpdateRequest
type TeacherUpdateRequest struct {
	// Id of the teacher
	// in: path
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	Id string `json:"id"`

	// in:body
	Body TeacherUpdateRequestBody `json:"body"`
}

type TeacherUpdateRequestBody struct {
	// Job Title Id
	// in: string
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	JobTitleId *string `json:"job_title_id,omitempty"`

	// Status of the teacher
	// in: string
	// example: CONTRACT
	Status *string `json:"status,omitempty"`

	// Experience of the teacher
	// in: int
	// example: 12
	Experience *int `json:"experience,omitempty"`

	// Degree of the teacher
	// in: string
	// example: B.Ed
	Degree *string `json:"degree,omitempty"`
}

// swagger:parameters TeacherDeleteRequest
type TeacherDeleteRequest struct {
	// Id of teacher
	// in: path
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	Id string `json:"id"`
}
