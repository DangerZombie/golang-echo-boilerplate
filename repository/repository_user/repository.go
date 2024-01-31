package repository_user

import (
	"go-echo/model/parameter"
	"go-echo/repository"

	"gorm.io/gorm"
)

type userRepo struct {
	base repository.BaseRepository
}

type UserRepository interface {
	FindUserById(db *gorm.DB, input parameter.FindUserByIdInput) (output parameter.FindUserByIdOutput, err error)
	FindUserByUsernameAndPassword(db *gorm.DB, input parameter.FindUserByUsernameAndPasswordInput) (output parameter.FindUserByUsernameAndPasswordOutput, err error)
}

func NewUserRepository(br repository.BaseRepository) UserRepository {
	return &userRepo{br}
}
