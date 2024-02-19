package service_teacher

import (
	"go-echo/helper/message"
	"go-echo/model/parameter"
	"go-echo/model/request"
	"go-echo/model/response"

	"go.uber.org/zap"
)

// swagger:operation GET /teacher/{id} Teacher TeacherDetailRequest
// Teacher Detail
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
//               $ref: '#/definitions/TeacherDetailResponse'
//           type: object

func (s *teacherServiceImpl) TeacherDetail(req request.TeacherDetailRequest) (res response.TeacherDetailResponse, msg message.Message, errMsg map[string]string) {
	logger := s.logger.With(zap.String("TeacherService", "TeacherDetail"))
	errMsg = map[string]string{}

	if req.Id == "" {
		logger.Error("log", zap.String("error", "id is empty"))
		errMsg["id"] = "id is empty"
		return res, message.FailedMsg, errMsg
	}

	tx := s.baseRepo.GetBegin()
	defer func() {
		if msg != message.SuccessMsg {
			s.baseRepo.BeginRollback(tx)
		} else {
			s.baseRepo.BeginCommit(tx)
		}
	}()

	findTeacherByIdInput := parameter.FindTeacherByIdInput{
		Id: req.Id,
	}

	teacher, err := s.teacherRepo.FindTeacherById(tx, findTeacherByIdInput)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["teacher"] = "error has been occured while looking teacher"
		return res, message.FailedMsg, errMsg
	}

	res = response.TeacherDetailResponse(teacher)

	return res, message.SuccessMsg, nil
}
