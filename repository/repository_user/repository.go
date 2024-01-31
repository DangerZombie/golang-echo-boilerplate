package repository_user

import (
	"go-echo/repository"
)

type userRepo struct {
	base repository.BaseRepository
}

func NewUserRepository(br repository.BaseRepository) UserRepository {
	return &userRepo{br}
}
