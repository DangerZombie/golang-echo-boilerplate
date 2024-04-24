package endpoint

import (
	"go-echo/service/service_teacher"
	"go-echo/service/service_user"

	"github.com/labstack/echo/v4"
)

type Endpoint interface {
	// Endpoint User
	LoginRequest(ctx echo.Context, s service_user.UserService) (int, interface{})
	RegisterUserRequest(ctx echo.Context, s service_user.UserService) (int, interface{})
	UserProfileRequest(ctx echo.Context, s service_user.UserService) (int, interface{})

	// Endpoint Teacher
	CreateTeacherRequest(ctx echo.Context, s service_teacher.TeacherService) (int, interface{})
	DeleteTeacherRequest(ctx echo.Context, s service_teacher.TeacherService) (int, interface{})
	FindTeacherDetailRequest(ctx echo.Context, s service_teacher.TeacherService) (int, interface{})
	ListTeachersRequest(ctx echo.Context, s service_teacher.TeacherService) (int, interface{})
	UpdateTeacherRequest(ctx echo.Context, s service_teacher.TeacherService) (int, interface{})
}
