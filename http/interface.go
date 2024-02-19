package http

import (
	"go-echo/service/service_teacher"
	"go-echo/service/service_user"

	"github.com/labstack/echo/v4"
)

type Http interface {
	// Swagger documentation
	SwaggerHttpHandler(r *echo.Echo)

	// API handler
	TeacherHandler(group *echo.Group, s service_teacher.TeacherService)
	UserHandler(group *echo.Group, s service_user.UserService)
}
