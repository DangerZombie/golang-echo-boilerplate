package parameter

import "go-echo/model/entity"

type CreateTeacherInput struct {
	entity.Teacher
}

type CreateTeacherOutput struct {
	Id string
}

type FindTeacherByIdInput struct {
	Id string
}

type FindTeacherByIdOutput struct {
	Id           string
	Nickname     string
	Email        string
	Status       string
	Experience   int
	Degree       string
	JobTitleId   string
	JobTitleName string
}

type UpdateTeacherByIdInput struct {
	Id     string
	Fields map[string]interface{}
}

type UpdateTeacherByIdOutput struct {
	entity.Teacher
}

type DeleteTeacherByIdInput struct {
	Id string
}

type DeleteTeacherByIdOutput struct {
	Success bool
}

type ListTeacherInput struct {
	Limit  int
	Page   int
	Sort   string
	Dir    string
	Filter map[string]interface{}
}

type ListTeacherOutput struct {
	Items []entity.Teacher
}
