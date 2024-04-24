package http

import (
	"go-echo/endpoint"
	"go-echo/service/service_teacher"

	"github.com/labstack/echo/v4"
)

func (h *httpImpl) TeacherHandler(group *echo.Group, s service_teacher.TeacherService) {
	group.POST("", func(ctx echo.Context) error {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).CreateTeacherRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})

	group.GET("", func(ctx echo.Context) error {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).ListTeachersRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})

	group.GET("/:id", func(ctx echo.Context) error {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).FindTeacherDetailRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})

	group.PUT("/:id", func(ctx echo.Context) error {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).UpdateTeacherRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})

	group.DELETE("/:id", func(ctx echo.Context) error {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).DeleteTeacherRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})
}
