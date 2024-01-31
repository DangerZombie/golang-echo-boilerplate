package repository_user

import (
	"errors"
	"go-echo/model/entity"
	"go-echo/model/parameter"

	"gorm.io/gorm"
)

func (r *userRepo) FindUserById(db *gorm.DB, input parameter.FindUserByIdInput) (output parameter.FindUserByIdOutput, err error) {
	err = db.
		Model(&entity.User{}).
		Where("id = ?", input.Id).
		First(&output).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}

		return output, err
	}

	return
}
