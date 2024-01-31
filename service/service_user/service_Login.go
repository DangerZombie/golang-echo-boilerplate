package service_user

import (
	"go-echo/helper/message"
	"go-echo/model/parameter"
	"go-echo/model/request"
	"go-echo/model/response"

	"go.uber.org/zap"
)

// swagger:operation GET /user/login User LoginRequest
// Login user
//
// ---
// responses:
//   '200':
//     description: Success Response.
//     schema:
//       properties:
//         meta:
//           $ref: '#/definitions/MetaSingleSuccessResponse'
//         data:
//           properties:
//             record:
//               $ref: '#/definitions/LoginResponse'
//           type: object

func (s *userServiceImpl) Login(req request.LoginRequest) (res response.LoginResponse, msg message.Message, errMsg map[string]string) {
	logger := s.logger.With(zap.String("UserService", "Login"))
	errMsg = map[string]string{}

	if req.Username == "" || req.Password == "" {
		logger.Error("log", zap.String("error", "field cannot be empty"))
		errMsg["user"] = "field cannot be empty"
		return res, message.FailedMsg, errMsg
	}

	tx := s.baseRepo.GetBegin()

	findUserByUsernameAndPasswordInput := parameter.FindUserByUsernameAndPasswordInput{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := s.userRepo.FindUserByUsernameAndPassword(tx, findUserByUsernameAndPasswordInput)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["user"] = "user invalid"
		return res, message.FailedMsg, errMsg
	}

	token, err := s.authHelper.GenerateJWT(user.Username)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["user"] = "error has been occured while generating token"
		return res, message.FailedMsg, errMsg
	}

	res = response.LoginResponse{
		Token: token,
	}

	return res, message.SuccessMsg, nil
}
