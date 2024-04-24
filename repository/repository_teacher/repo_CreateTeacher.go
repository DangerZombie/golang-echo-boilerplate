package repository_teacher

import (
	"go-echo/model/parameter"

	"gorm.io/gorm"
)

func (r *teacherRepo) CreateTeacher(db *gorm.DB, input parameter.CreateTeacherInput) (output parameter.CreateTeacherOutput, err error) {
	err = db.Debug().Create(&input.Teacher).Error
	if err != nil {
		return output, err
	}

	output.Id = input.Id

	return output, nil
}
