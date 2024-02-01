package service_user

import (
	"go-echo/helper/message"
	"go-echo/model/request"
	"go-echo/model/response"
)

type UserService interface {
	Login(req request.LoginRequest) (res response.LoginResponse, msg message.Message, errMsg map[string]string)
	UserProfile(req request.UserProfileRequest) (res response.UserProfileResponse, msg message.Message, errMsg map[string]string)
}
