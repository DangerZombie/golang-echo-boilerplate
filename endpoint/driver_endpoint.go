package endpoint

import (
	"encoding/json"
	"go-echo/helper/auth"
	"go-echo/helper/message"
	"go-echo/model/base"
	"go-echo/model/request"
	"go-echo/service"
	"net/http"

	"github.com/labstack/echo"
)

func InsertDriverRequest(ctx echo.Context, s service.DriverService) interface{} {
	req := request.InsertDriverRequest{}
	_ = json.NewDecoder(ctx.Request().Body).Decode(&req)
	result, msg, errMsg := s.InsertDriver(req)

	var wrap interface{}
	if msg.Code == 4000 {
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
	} else {
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, nil, errMsg)
	}

	return wrap
}

func GetListDriversRequest(ctx echo.Context, s service.DriverService) interface{} {
	req := request.GetListDriversRequest{}
	_ = json.NewDecoder(ctx.Request().Body).Decode(&req)
	result, pagination, msg, errMsg := s.GetListDrivers(req)

	var wrap interface{}
	if msg.Code == 4000 {
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
	} else {
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, pagination, errMsg)
	}

	return wrap
}

func GetDriverByNumberRequest(ctx echo.Context, s service.DriverService) (int, interface{}) {
	// JWT verify
	var wrap interface{}
	var status int
	_, err := auth.VerifyJWT(ctx.Request().Header)
	if err != nil {
		wrap = base.SetHttpResponse(message.ErrNoAuth.Code, message.ErrNoAuth.Message, nil, nil, nil)
		status = http.StatusUnauthorized
		return status, wrap
	}

	req := request.GetDriverByNumber{}
	req.Number = ctx.Param("number")
	result, msg, errMsg := s.GetDriverByNumber(req)

	if msg.Code == 4000 {
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
		status = http.StatusInternalServerError
	} else {
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, nil, errMsg)
		status = http.StatusOK
	}

	return status, wrap
}

func UpdateDriverByNumberRequest(ctx echo.Context, s service.DriverService) interface{} {
	req := request.UpdateDriverByNumber{}
	_ = json.NewDecoder(ctx.Request().Body).Decode(&req)
	result, msg, errMsg := s.UpdateDriverByNumber(req)

	var wrap interface{}
	if msg.Code == 4000 {
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
	} else {
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, nil, errMsg)
	}

	return wrap
}

func DeleteDriverByNumberRequest(ctx echo.Context, s service.DriverService) interface{} {
	req := request.DeleteDriverByNumber{}
	_ = json.NewDecoder(ctx.Request().Body).Decode(&req)
	result, msg, errMsg := s.DeleteDriverByNumber(req)

	var wrap interface{}
	if msg.Code == 4000 {
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
	} else {
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, nil, errMsg)
	}

	return wrap
}
