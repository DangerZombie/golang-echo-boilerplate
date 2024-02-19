package repository_teacher

import (
	"go-echo/model/entity"
	"go-echo/model/parameter"

	"gorm.io/gorm"
)

func (r *teacherRepo) DeleteTeacherById(db *gorm.DB, input parameter.DeleteTeacherByIdInput) (output parameter.DeleteTeacherByIdOutput, err error) {
	err = db.
		Where("id = ?", input.Id).
		Delete(&entity.Teacher{}).
		Error

	if err != nil {
		return
	}

	output.Success = true

	return
}
