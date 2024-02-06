package repository_teacher

import (
	"go-echo/model/base"
	"go-echo/model/parameter"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	CreateTeacher(db *gorm.DB, input parameter.CreateTeacherInput) (output parameter.CreateTeacherOutput, err error)
	DeleteTeacherById(db *gorm.DB, input parameter.DeleteTeacherByIdInput) (output parameter.DeleteTeacherByIdOutput, err error)
	FindTeacherById(db *gorm.DB, input parameter.FindTeacherByIdInput) (output parameter.FindTeacherByIdOutput, err error)
	ListTeacher(db *gorm.DB, input parameter.ListTeacherInput) (output parameter.ListTeacherOutput, pagination base.Pagination, err error)
	UpdateTeacherById(db *gorm.DB, input parameter.UpdateTeacherByIdInput) (output parameter.UpdateTeacherByIdOutput, err error)
}
