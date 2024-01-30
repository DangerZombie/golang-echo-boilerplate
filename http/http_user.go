package http

import (
	"go-echo/endpoint"
	"go-echo/service/service_user"
	"net/http"

	"github.com/labstack/echo"
)

func UserHandler(group *echo.Group, s service_user.UserService) {
	group.GET("/login", func(ctx echo.Context) error {
		result := endpoint.LoginRequest(ctx, s)
		return ctx.JSON(http.StatusOK, result)
	})
}
