package service_user

import (
	"go-echo/helper/message"
	"go-echo/model/entity"
	"go-echo/model/parameter"
	"go-echo/model/request"
	"go-echo/model/response"

	"go.uber.org/zap"
)

// swagger:operation POST /user/register User RegisterUserRequest
// Register user
//
// ---
// security:
// - Bearer: []
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
//               $ref: '#/definitions/RegisterUserResponse'
//           type: object

func (s *userServiceImpl) RegisterUser(req request.RegisterUserRequestBody) (res response.RegisterUserResponse, msg message.Message, errMsg map[string]string) {
	logger := s.logger.With(zap.String("UserService", "RegisterUser"))
	errMsg = map[string]string{}

	tx := s.baseRepo.GetBegin()
	defer func() {
		if msg != message.SuccessMsg {
			s.baseRepo.BeginRollback(tx)
		} else {
			s.baseRepo.BeginCommit(tx)
		}
	}()

	createUserInput := parameter.CreateUserInput{
		User: entity.User{
			Username:  req.Username,
			Password:  req.Password,
			Nickname:  req.Nickname,
			Status:    "ACTIVE",
			CreatedBy: "system",
		},
	}

	user, err := s.userRepo.CreateUser(tx, createUserInput)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["user"] = "error has been occured while inserting user"
		return res, message.FailedMsg, errMsg
	}

	res = response.RegisterUserResponse{
		Id: user.Id,
	}

	return res, message.SuccessMsg, nil
}
