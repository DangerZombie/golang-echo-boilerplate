package user

import (
	"errors"
	"go-echo/model/entity"

	"gorm.io/gorm"
)

func (r *userRepo) GetUser(db *gorm.DB, username string, password string) (*entity.User, error) {
	var user entity.User
	err := db.
		Model(&entity.User{}).
		Where("username = ? AND password = ?", username, password).
		First(&user).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
