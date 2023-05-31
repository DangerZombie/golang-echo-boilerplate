package http

import (
	"go-echo/endpoint"
	"go-echo/service"
	"net/http"

	"github.com/labstack/echo"
)

func UserHandler(group *echo.Group, s service.UserService) {
	group.POST("/", func(ctx echo.Context) error {
		result := endpoint.LoginRequest(ctx, s)
		return ctx.JSON(http.StatusOK, result)
	})
}
