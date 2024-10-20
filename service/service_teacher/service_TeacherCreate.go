package service_teacher

import (
	"go-echo/helper/message"
	"go-echo/model/base"
	"go-echo/model/entity"
	"go-echo/model/parameter"
	"go-echo/model/request"
	"go-echo/model/response"

	"go.uber.org/zap"
)

// swagger:operation POST /teacher Teacher TeacherCreateRequest
// Add Teacher
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
//               $ref: '#/definitions/TeacherCreateResponse'
//           type: object

func (s *teacherServiceImpl) TeacherCreate(req request.TeacherCreateRequestBody) (res response.TeacherCreateResponse, msg message.Message, errMsg map[string]string) {
	logger := s.logger.With(zap.String("TeacherService", "TeacherCreate"))
	errMsg = map[string]string{}

	teacherInput := parameter.CreateTeacherInput{
		Teacher: entity.Teacher{
			UserId:     req.UserId,
			JobTitleId: req.JobTitleId,
			Status:     req.Status,
			Experience: req.Experience,
			Degree:     req.Degree,
			BaseModel: base.BaseModel{
				CreatedBy: req.Issuer,
				UpdatedBy: req.Issuer,
			},
		},
	}

	tx := s.baseRepo.GetBegin()
	defer func() {
		if msg != message.SuccessMsg {
			s.baseRepo.BeginRollback(tx)
		} else {
			s.baseRepo.BeginCommit(tx)
		}
	}()

	teacher, err := s.teacherRepo.CreateTeacher(tx, teacherInput)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["teacher"] = "error has been occured while inserting teacher"
		return res, message.FailedMsg, errMsg
	}

	res = response.TeacherCreateResponse{
		Id: teacher.Id,
	}

	return res, message.SuccessMsg, nil
}
