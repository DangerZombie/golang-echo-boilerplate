package service_teacher

import (
	"go-echo/helper/message"
	"go-echo/model/parameter"
	"go-echo/model/request"
	"go-echo/model/response"

	"go.uber.org/zap"
)

// swagger:operation PUT /teacher/{id} Teacher TeacherUpdateRequest
// Teacher Update
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
//               $ref: '#/definitions/TeacherUpdateResponse'
//           type: object

func (s *teacherServiceImpl) TeacherUpdate(req request.TeacherUpdateRequest) (res response.TeacherUpdateResponse, msg message.Message, errMsg map[string]string) {
	logger := s.logger.With(zap.String("TeacherService", "TeacherUpdate"))
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

	updateTeacherByIdInput := parameter.UpdateTeacherByIdInput{
		Id: req.Id,
		Fields: map[string]interface{}{
			"job_title_id": &req.Body.JobTitleId,
			"status":       &req.Body.Status,
			"experience":   &req.Body.Experience,
			"degree":       &req.Body.Degree,
		},
	}

	teacher, err := s.teacherRepo.UpdateTeacherById(tx, updateTeacherByIdInput)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["teacher"] = "error has been occured while updating teacher"
		return res, message.FailedMsg, errMsg
	}

	res = response.TeacherUpdateResponse{
		Id:         teacher.Id,
		JobTitleId: teacher.JobTitleId,
		Status:     teacher.Status,
		Experience: teacher.Experience,
		Degree:     teacher.Degree,
	}

	return res, message.SuccessMsg, nil
}
