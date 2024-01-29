package service

import (
	"go-echo/helper/auth"
	"go-echo/helper/message"
	"go-echo/model/request"
	"go-echo/model/response"
	"go-echo/repository"
	"go-echo/repository/user"

	"go.uber.org/zap"
)

type userServiceImpl struct {
	logger   *zap.Logger
	baseRepo repository.BaseRepository
	userRepo user.UserRepository
}

type UserService interface {
	Login(req request.LoginRequest) (*response.LoginResponse, message.Message, interface{})
}

func NewUserService(
	lg *zap.Logger,
	br repository.BaseRepository,
	ur user.UserRepository,
) UserService {
	return &userServiceImpl{lg, br, ur}
}

func (s *userServiceImpl) Login(req request.LoginRequest) (*response.LoginResponse, message.Message, interface{}) {
	logger := s.logger.With(zap.String("UserService", "Login"))
	errMsg := map[string]string{}

	if req.Username == "" || req.Password == "" {
		logger.Error("log", zap.String("error", "field cannot be empty"))
		errMsg["user"] = "field cannot be empty"
		return nil, message.FailedMsg, errMsg
	}

	tx := s.baseRepo.GetBegin()
	user, err := s.userRepo.GetUser(tx, req.Username, req.Password)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["user"] = "user invalid"
		return nil, message.FailedMsg, errMsg
	}

	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["user"] = "error has been occured while generating token"
		return nil, message.FailedMsg, errMsg
	}

	result := response.LoginResponse{
		Token: token,
	}

	return &result, message.SuccessMsg, nil
}
