package service_user

import (
	"go-echo/helper/auth"
	"go-echo/repository"
	"go-echo/repository/repository_user"

	"go.uber.org/zap"
)

type userServiceImpl struct {
	logger     *zap.Logger
	authHelper auth.AuthHelper
	baseRepo   repository.BaseRepository
	userRepo   repository_user.UserRepository
}

func NewUserService(
	lg *zap.Logger,
	ah auth.AuthHelper,
	br repository.BaseRepository,
	ur repository_user.UserRepository,
) UserService {
	return &userServiceImpl{lg, ah, br, ur}
}
