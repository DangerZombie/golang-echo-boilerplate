package response

// swagger:model TeacherCreateResponse
type TeacherCreateResponse struct {
	// Id of the teacher
	// in: string
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	Id string `json:"id"`
}

// swagger:model TeacherListResponse
type TeacherListResponse struct {
	// Id of teacher
	// in: string
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	Id string `json:"id"`

	// Name of teacher
	// in: string
	// example: John Doe
	Nickname string `json:"nickname"`

	// Email of teacher
	// in: string
	// example: teacher@school.edu
	Email string `json:"email"`

	// Status of the teacher
	// in: string
	// example: CONTRACT
	Status string `json:"status"`

	// Experience of the teacher
	// in: int
	// example: 12
	Experience int `json:"experience"`

	// Last Education with degree
	// in: string
	// example: B.Ed
	Degree string `json:"degree"`
}

// swagger:model TeacherDetailResponse
type TeacherDetailResponse struct {
	// Id of the teacher
	// in: string
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	Id string `json:"id"`

	// Nickname is the name of the teacher
	// in: string
	// example: John
	Nickname string `json:"nickname"`

	// Email of the teacher
	// in: string
	// example: teacher@school.edu
	Email string `json:"email"`

	// Status of the teacher
	// in: string
	// example: CONTRACT
	Status string `json:"status"`

	// Experience of the teacher
	// in: int
	// example: 12
	Experience int `json:"experience"`

	// Last Education with degree
	// in: string
	// example: B.Ed
	Degree string `json:"degree"`

	// Job Title Id of the teacher
	// in: string
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	JobTitleId string `json:"job_title_id"`

	// Job Title Name of the teacher
	// in: string
	// example: Principal
	JobTitleName string `json:"job_title_name"`
}

// swagger:model TeacherUpdateResponse
type TeacherUpdateResponse struct {
	// Id of the teacher
	// in: string
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	Id string `json:"id"`

	// Job Title Id of the teacher
	// in: string
	// example: 4fc427da-91c7-45b5-b4f9-f6dcc646005f
	JobTitleId string `json:"job_title_Id"`

	// Status of the teacher
	// in: string
	// example: CONTRACT
	Status string `json:"status"`

	// Experience of the teacher
	// in: int
	// example: 12
	Experience int `json:"experience"`

	// Degree of the teacher
	// in: string
	// example: B.Ed
	Degree string `json:"degree"`
}

// swagger:model TeacherDeleteResponse
type TeacherDeleteResponse struct {
	// Success message of deleted data
	// in: bool
	// example: true
	Success bool `json:"success"`
}
