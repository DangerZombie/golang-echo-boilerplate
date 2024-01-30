package repository_user

import (
	"go-echo/model/entity"
	"go-echo/repository"

	"gorm.io/gorm"
)

type userRepo struct {
	base repository.BaseRepository
}

type UserRepository interface {
	FindUserByUsernameAndPassword(db *gorm.DB, username string, password string) (*entity.User, error)
}

func NewUserRepository(br repository.BaseRepository) UserRepository {
	return &userRepo{br}
}
