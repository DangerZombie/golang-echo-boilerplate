package service_teacher

import (
	"go-echo/helper/message"
	"go-echo/model/parameter"
	"go-echo/model/request"
	"go-echo/model/response"

	"go.uber.org/zap"
)

// swagger:operation DELETE /teacher/{id} Teacher TeacherDeleteRequest
// Delete Teacher
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
//               $ref: '#/definitions/TeacherDeleteResponse'
//           type: object

func (s *teacherServiceImpl) TeacherDelete(req request.TeacherDeleteRequest) (res response.TeacherDeleteResponse, msg message.Message, errMsg map[string]string) {
	logger := s.logger.With(zap.String("TeacherService", "TeacherDelete"))
	errMsg = map[string]string{}

	if req.Id == "" {
		logger.Error("log", zap.String("error", "id is empty"))
		errMsg["teacher"] = "id is empty"
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

	deleteTeacherInput := parameter.DeleteTeacherByIdInput{
		Id: req.Id,
	}

	output, err := s.teacherRepo.DeleteTeacherById(tx, deleteTeacherInput)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["teacher"] = "error has been occured while deleting teacher"
		return res, message.FailedMsg, errMsg
	}

	res = response.TeacherDeleteResponse{
		Success: output.Success,
	}

	return res, message.SuccessMsg, nil
}
