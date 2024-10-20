package service_teacher

import (
	"go-echo/helper/message"
	"go-echo/model/base"
	"go-echo/model/parameter"
	"go-echo/model/request"
	"go-echo/model/response"

	"go.uber.org/zap"
)

// swagger:operation GET /teacher Teacher TeacherListRequest
// Teacher List
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
//               $ref: '#/definitions/TeacherListResponse'
//           type: object

func (s *teacherServiceImpl) TeacherList(req request.TeacherListRequest) (res []response.TeacherListResponse, pagination base.Pagination, msg message.Message, errMsg map[string]string) {
	logger := s.logger.With(zap.String("TeacherService", "TeacerList"))
	errMsg = map[string]string{}

	if req.Page < 1 {
		req.Page = 1
	}

	if req.Limit < 1 {
		req.Limit = 10
	}

	filter := map[string]interface{}{
		"name": req.Name,
	}

	tx := s.baseRepo.GetBegin()
	defer func() {
		if msg != message.SuccessMsg {
			s.baseRepo.BeginRollback(tx)
		} else {
			s.baseRepo.BeginCommit(tx)
		}
	}()

	teacherListInput := parameter.ListTeacherInput{
		Limit:  req.Limit,
		Page:   req.Page,
		Sort:   req.Sort,
		Dir:    req.Dir,
		Filter: filter,
	}

	teachers, pagination, err := s.teacherRepo.ListTeacher(tx, teacherListInput)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["teacher"] = "error has been occured while looking teacher"
		return res, pagination, message.FailedMsg, errMsg
	}

	for _, val := range teachers.Items {
		teacher := response.TeacherListResponse{
			Id:         val.Id,
			Nickname:   val.User.Nickname,
			Email:      val.User.Username,
			Status:     val.Status,
			Experience: val.Experience,
			Degree:     val.Degree,
		}

		res = append(res, teacher)
	}

	return res, pagination, message.SuccessMsg, nil
}
