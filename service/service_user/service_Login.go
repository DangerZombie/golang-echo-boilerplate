package service_user

import (
	"go-echo/helper/message"
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

func (s *userServiceImpl) Login(req request.LoginRequest) (*response.LoginResponse, message.Message, interface{}) {
	logger := s.logger.With(zap.String("UserService", "Login"))
	errMsg := map[string]string{}

	if req.Username == "" || req.Password == "" {
		logger.Error("log", zap.String("error", "field cannot be empty"))
		errMsg["user"] = "field cannot be empty"
		return nil, message.FailedMsg, errMsg
	}

	tx := s.baseRepo.GetBegin()
	user, err := s.userRepo.FindUserByUsernameAndPassword(tx, req.Username, req.Password)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["user"] = "user invalid"
		return nil, message.FailedMsg, errMsg
	}

	token, err := s.authHelper.GenerateJWT(user.Username)
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
