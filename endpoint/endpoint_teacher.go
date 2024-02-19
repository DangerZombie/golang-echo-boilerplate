package endpoint

import (
	"encoding/json"
	"go-echo/helper/message"
	"go-echo/helper/static"
	"go-echo/helper/util"
	"go-echo/model/base"
	"go-echo/model/request"
	"go-echo/service/service_teacher"
	"net/http"

	"github.com/go-faker/faker/v4/pkg/slice"
	"github.com/labstack/echo/v4"
)

func (e *endpointImpl) CreateTeacherRequest(ctx echo.Context, s service_teacher.TeacherService) (int, interface{}) {
	// Verify JWT token from the request headers
	claims, err := e.authHelper.VerifyJWT(ctx.Request().Header)
	if err != nil {
		wrap := base.SetHttpResponse(message.ErrNoAuth.Code, message.ErrNoAuth.Message, nil, nil, map[string]string{"token": err.Error()})
		return http.StatusUnauthorized, wrap
	}

	// Validate allowed roles
	if !slice.Contains(claims.Roles, static.RoleADMINISTRATOR) {
		wrap := base.SetHttpResponse(message.ErrForbiddenAccess.Code, message.ErrForbiddenAccess.Message, nil, nil, map[string]string{"role": "you not have correct role"})
		return http.StatusForbidden, wrap
	}

	req := request.TeacherCreateRequestBody{}
	_ = json.NewDecoder(ctx.Request().Body).Decode(&req)
	req.Issuer = claims.Issuer
	result, msg, errMsg := s.TeacherCreate(req)

	var wrap interface{}
	var code int
	if msg.Code == 4000 {
		code = http.StatusBadRequest
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
	} else {
		code = http.StatusOK
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, nil, errMsg)
	}

	return code, wrap
}

func (e *endpointImpl) ListTeachersRequest(ctx echo.Context, s service_teacher.TeacherService) (int, interface{}) {
	// Verify JWT token from the request headers
	_, err := e.authHelper.VerifyJWT(ctx.Request().Header)
	if err != nil {
		wrap := base.SetHttpResponse(message.ErrNoAuth.Code, message.ErrNoAuth.Message, nil, nil, map[string]string{"token": err.Error()})
		return http.StatusUnauthorized, wrap
	}

	req := request.TeacherListRequest{
		Page:  util.StringToInt(ctx.QueryParam("page")),
		Limit: util.StringToInt(ctx.QueryParam("limit")),
		Sort:  ctx.QueryParam("sort"),
		Dir:   ctx.QueryParam("dir"),
		Name:  ctx.QueryParam("name"),
	}

	result, pagination, msg, errMsg := s.TeacherList(req)

	var wrap interface{}
	var code int
	if msg.Code == 4000 {
		code = http.StatusBadRequest
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
	} else {
		code = http.StatusOK
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, &pagination, errMsg)
	}

	return code, wrap
}

func (e *endpointImpl) FindTeacherDetailRequest(ctx echo.Context, s service_teacher.TeacherService) (int, interface{}) {
	// Verify JWT token from the request headers
	_, err := e.authHelper.VerifyJWT(ctx.Request().Header)
	if err != nil {
		wrap := base.SetHttpResponse(message.ErrNoAuth.Code, message.ErrNoAuth.Message, nil, nil, map[string]string{"token": err.Error()})
		return http.StatusUnauthorized, wrap
	}

	req := request.TeacherDetailRequest{}
	req.Id = ctx.Param("id")
	result, msg, errMsg := s.TeacherDetail(req)

	var wrap interface{}
	var code int
	if msg.Code == 4000 {
		code = http.StatusBadRequest
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
	} else {
		code = http.StatusOK
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, nil, errMsg)
	}

	return code, wrap
}

func (e *endpointImpl) UpdateTeacherRequest(ctx echo.Context, s service_teacher.TeacherService) (int, interface{}) {
	// Verify JWT token from the request headers
	_, err := e.authHelper.VerifyJWT(ctx.Request().Header)
	if err != nil {
		wrap := base.SetHttpResponse(message.ErrNoAuth.Code, message.ErrNoAuth.Message, nil, nil, map[string]string{"token": err.Error()})
		return http.StatusUnauthorized, wrap
	}

	reqBody := request.TeacherUpdateRequestBody{}
	_ = json.NewDecoder(ctx.Request().Body).Decode(&reqBody)
	req := request.TeacherUpdateRequest{
		Id:   ctx.Param("id"),
		Body: reqBody,
	}

	result, msg, errMsg := s.TeacherUpdate(req)

	var wrap interface{}
	var code int
	if msg.Code == 4000 {
		code = http.StatusBadRequest
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
	} else {
		code = http.StatusOK
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, nil, errMsg)
	}

	return code, wrap
}

func (e *endpointImpl) DeleteTeacherRequest(ctx echo.Context, s service_teacher.TeacherService) (int, interface{}) {
	// Verify JWT token from the request headers
	_, err := e.authHelper.VerifyJWT(ctx.Request().Header)
	if err != nil {
		wrap := base.SetHttpResponse(message.ErrNoAuth.Code, message.ErrNoAuth.Message, nil, nil, map[string]string{"token": err.Error()})
		return http.StatusUnauthorized, wrap
	}

	req := request.TeacherDeleteRequest{}
	req.Id = ctx.Param("id")
	result, msg, errMsg := s.TeacherDelete(req)

	var wrap interface{}
	var code int
	if msg.Code == 4000 {
		code = http.StatusBadRequest
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
	} else {
		code = http.StatusOK
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, nil, errMsg)
	}

	return code, wrap
}
