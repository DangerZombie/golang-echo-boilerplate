package endpoint

import (
	"go-echo/helper/auth"
	"go-echo/service/service_user"

	"github.com/labstack/echo"
)

type endpointImpl struct {
	authHelper auth.AuthHelper
}

type Endpoint interface {
	// Endpoint User
	LoginRequest(ctx echo.Context, s service_user.UserService) (int, interface{})
	UserProfileRequest(ctx echo.Context, s service_user.UserService) (int, interface{})
}

func NewEndpoint(ah auth.AuthHelper) Endpoint {
	return &endpointImpl{ah}
}
