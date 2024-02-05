package auth

import (
	"go-echo/repository"
	"go-echo/repository/repository_user"
)

type authHelperImpl struct {
	baseRepo repository.BaseRepository
	userRepo repository_user.UserRepository
}

func NewAuthHelper(br repository.BaseRepository, ur repository_user.UserRepository) AuthHelper {
	return &authHelperImpl{br, ur}
}
