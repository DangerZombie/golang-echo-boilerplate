package http

import (
	"go-echo/helper/auth"
	"go-echo/service/service_user"

	"github.com/labstack/echo"
)

type httpImpl struct {
	authHelper auth.AuthHelper
}

type Http interface {
	UserHandler(group *echo.Group, s service_user.UserService)
}

func NewHttp(ah auth.AuthHelper) Http {
	return &httpImpl{ah}
}
