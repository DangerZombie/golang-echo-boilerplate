package endpoint

import (
	"encoding/json"
	"go-echo/model/base"
	"go-echo/model/request"
	"go-echo/service"

	"github.com/labstack/echo"
)

func LoginRequest(ctx echo.Context, s service.UserService) interface{} {
	req := request.LoginRequest{}
	_ = json.NewDecoder(ctx.Request().Body).Decode(&req)
	result, msg, errMsg := s.Login(req)

	var wrap interface{}
	if msg.Code == 4000 {
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
	} else {
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, nil, errMsg)
	}

	return wrap
}
