package http

import (
	"go-echo/endpoint"
	"go-echo/service/service_driver"
	"net/http"

	"github.com/labstack/echo"
)

func DriverHandler(group *echo.Group, s service_driver.DriverService) {
	group.POST("/", func(ctx echo.Context) error {
		result := endpoint.InsertDriverRequest(ctx, s)
		return ctx.JSON(http.StatusOK, result)
	})

	group.GET("/", func(ctx echo.Context) error {
		result := endpoint.GetListDriversRequest(ctx, s)
		return ctx.JSON(http.StatusOK, result)
	})

	group.GET("/:number", func(ctx echo.Context) error {
		statusCode, result := endpoint.GetDriverByNumberRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})

	group.PUT("/:number", func(ctx echo.Context) error {
		result := endpoint.UpdateDriverByNumberRequest(ctx, s)
		return ctx.JSON(http.StatusOK, result)
	})

	group.DELETE("/:number", func(ctx echo.Context) error {
		result := endpoint.DeleteDriverByNumberRequest(ctx, s)
		return ctx.JSON(http.StatusOK, result)
	})
}
