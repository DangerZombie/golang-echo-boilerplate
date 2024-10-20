package http

import (
	"go-echo/endpoint"
	"go-echo/service/service_user"

	"github.com/labstack/echo/v4"
)

func (h *httpImpl) UserHandler(group *echo.Group, s service_user.UserService) {
	group.POST("/login", func(ctx echo.Context) error {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).LoginRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})

	group.GET("/profile", func(ctx echo.Context) error {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).UserProfileRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})

	group.POST("/register", func(ctx echo.Context) error {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).RegisterUserRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})
}
