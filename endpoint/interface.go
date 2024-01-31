package endpoint

import (
	"go-echo/service/service_user"

	"github.com/labstack/echo"
)

type Endpoint interface {
	// Endpoint User
	LoginRequest(ctx echo.Context, s service_user.UserService) (int, interface{})
	UserProfileRequest(ctx echo.Context, s service_user.UserService) (int, interface{})
}
