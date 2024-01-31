package http

import (
	"go-echo/endpoint"
	"go-echo/service/service_user"

	"github.com/labstack/echo"
)

func UserHandler(group *echo.Group, s service_user.UserService) {
	group.GET("/login", func(ctx echo.Context) error {
		statusCode, result := endpoint.LoginRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})

	group.GET("/profile", func(ctx echo.Context) error {
		statusCode, result := endpoint.UserProfileRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})
}
