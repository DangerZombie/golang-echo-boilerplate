package http

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/labstack/echo"
)

func SwaggerHttpHandler(r *echo.Echo) {
    // serve swagger spec file
    r.GET("/swagger.yaml", func(c echo.Context) error {
        return c.File("swagger.yaml")
    })

    // serve swagger UI
    opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
    sh := middleware.SwaggerUI(opts, nil)
    r.GET("/docs/*", echo.WrapHandler(sh))

    // serve Redoc
    opts1 := middleware.RedocOpts{SpecURL: "/swagger.yaml", Path: "doc"}
    sh1 := middleware.Redoc(opts1, nil)
    r.GET("/doc/*", echo.WrapHandler(sh1))
}
