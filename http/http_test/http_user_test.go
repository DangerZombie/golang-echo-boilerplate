package http_test

import (
	"go-echo/helper/auth"
	transport "go-echo/http"
	"go-echo/service/service_user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

func TestHttpUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)
	mockUserService := service_user.NewMockUserService(mockCtrl)

	httpModule := transport.NewHttp(
		mockAuthHelper,
	)

	e := echo.New()
	g := e.Group("/api/v1/user")

	t.Run("Should create teacher", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		rec := httptest.NewRecorder()
		_ = e.NewContext(req, rec)

		httpModule.UserHandler(g, mockUserService)
	})
}
