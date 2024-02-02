package endpoint

import (
	"encoding/json"
	"go-echo/helper/message"
	"go-echo/model/base"
	"go-echo/model/request"
	"go-echo/service/service_user"
	"net/http"

	"github.com/labstack/echo"
)

func (e *endpointImpl) LoginRequest(ctx echo.Context, s service_user.UserService) (int, interface{}) {
	req := request.LoginRequest{}
	_ = json.NewDecoder(ctx.Request().Body).Decode(&req)
	result, msg, errMsg := s.Login(req)

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

func (e *endpointImpl) UserProfileRequest(ctx echo.Context, s service_user.UserService) (int, interface{}) {
	// Verify JWT token from the request headers
	_, err := e.authHelper.VerifyJWT(ctx.Request().Header)
	if err != nil {
		wrap := base.SetHttpResponse(message.ErrNoAuth.Code, message.ErrNoAuth.Message, nil, nil, map[string]string{"token": err.Error()})
		return http.StatusUnauthorized, wrap
	}

	userProfileInput := request.UserProfileRequest{
		Id: ctx.QueryParam("id"),
	}

	result, msg, errMsg := s.UserProfile(userProfileInput)

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

func (e *endpointImpl) RegisterUserRequest(ctx echo.Context, s service_user.UserService) (int, interface{}) {
	// Verify JWT token from the request headers
	_, err := e.authHelper.VerifyJWT(ctx.Request().Header)
	if err != nil {
		wrap := base.SetHttpResponse(message.ErrNoAuth.Code, message.ErrNoAuth.Message, nil, nil, map[string]string{"token": err.Error()})
		return http.StatusUnauthorized, wrap
	}

	req := request.RegisterUserRequest{}
	_ = json.NewDecoder(ctx.Request().Body).Decode(&req)

	registerUserInput := request.RegisterUserRequest{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
	}

	result, msg, errMsg := s.RegisterUser(registerUserInput)

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
