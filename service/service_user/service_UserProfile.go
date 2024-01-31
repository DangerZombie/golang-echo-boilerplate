package service_user

import (
	"go-echo/helper/message"
	"go-echo/model/parameter"
	"go-echo/model/request"
	"go-echo/model/response"

	"go.uber.org/zap"
)

func (s *userServiceImpl) UserProfile(req request.UserProfileRequest) (res response.UserProfileResponse, msg message.Message, errMsg map[string]string) {
	logger := s.logger.With(zap.String("UserService", "Login"))
	errMsg = map[string]string{}

	if req.Id == "" {
		logger.Error("log", zap.String("error", "field cannot be empty"))
		errMsg["id"] = "field cannot be empty"
		return res, message.FailedMsg, errMsg
	}

	tx := s.baseRepo.GetBegin()

	findUserByIdInput := parameter.FindUserByIdInput{
		Id: req.Id,
	}

	user, err := s.userRepo.FindUserById(tx, findUserByIdInput)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["id"] = "user id invalid"
		return res, message.FailedMsg, errMsg
	}

	res = response.UserProfileResponse{
		Id:       user.Id,
		Nickname: user.Nickname,
	}

	return res, message.SuccessMsg, nil
}
