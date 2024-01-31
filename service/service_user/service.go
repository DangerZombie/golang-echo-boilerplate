package service_user

import (
	"go-echo/helper/auth"
	"go-echo/helper/message"
	"go-echo/model/request"
	"go-echo/model/response"
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

type UserService interface {
	Login(req request.LoginRequest) (res response.LoginResponse, msg message.Message, errMsg map[string]string)
	UserProfile(req request.UserProfileRequest) (res response.UserProfileResponse, msg message.Message, errMsg map[string]string)
}

func NewUserService(
	lg *zap.Logger,
	ah auth.AuthHelper,
	br repository.BaseRepository,
	ur repository_user.UserRepository,
) UserService {
	return &userServiceImpl{lg, ah, br, ur}
}
