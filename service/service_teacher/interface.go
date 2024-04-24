package service_teacher

import (
	"go-echo/helper/message"
	"go-echo/model/base"
	"go-echo/model/request"
	"go-echo/model/response"
)

type TeacherService interface {
	TeacherCreate(req request.TeacherCreateRequestBody) (res response.TeacherCreateResponse, msg message.Message, errMsg map[string]string)
	TeacherDelete(req request.TeacherDeleteRequest) (res response.TeacherDeleteResponse, msg message.Message, errMsg map[string]string)
	TeacherDetail(req request.TeacherDetailRequest) (res response.TeacherDetailResponse, msg message.Message, errMsg map[string]string)
	TeacherList(req request.TeacherListRequest) (res []response.TeacherListResponse, pagination base.Pagination, msg message.Message, errMsg map[string]string)
	TeacherUpdate(req request.TeacherUpdateRequest) (res response.TeacherUpdateResponse, msg message.Message, errMsg map[string]string)
}
