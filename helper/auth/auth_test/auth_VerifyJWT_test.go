package auth_test

import (
	"go-echo/helper/auth"
	"go-echo/helper/static"
	"go-echo/model/base"
	"go-echo/model/entity"
	"go-echo/model/parameter"
	"go-echo/repository"
	"go-echo/repository/repository_user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestVerifyJWT(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockBaseRepository := repository.NewMockBaseRepository(mockCtrl)
	mockUserRepository := repository_user.NewMockUserRepository(mockCtrl)

	authHelper := auth.NewAuthHelper(
		mockBaseRepository,
		mockUserRepository,
	)

	e := echo.New()
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath("./../../..")
	// viper.SetConfigName("config-dev")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic("config not found")
	// }
	// viper.AutomaticEnv()

	id := faker.UUIDHyphenated()
	findUserRoleByUserIdInput := parameter.FindUserRoleByUserIdInput{
		Id: id,
	}

	findUserRoleByUserIdOutput := parameter.FindUserRoleByUserIdOutput{
		BaseModel: base.BaseModel{
			Id: id,
		},
		Username: faker.Username(),
		Password: faker.Password(),
		Status:   "ACTIVE",
		Nickname: faker.Name(),
		Roles: []*entity.Role{
			{
				Id:   faker.UUIDHyphenated(),
				Name: static.RoleADMINISTRATOR,
			},
		},
	}

	t.Run("Should return claims", func(t *testing.T) {
		// generate valid JWT
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserRoleByUserId(gomock.Any(), findUserRoleByUserIdInput).
			Times(1).
			Return(findUserRoleByUserIdOutput, nil)

		jwt, err := authHelper.GenerateJWT(id)

		require.NotEmpty(t, jwt)
		require.Empty(t, err)

		// Verify JWT
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+jwt)
		rec := httptest.NewRecorder()
		_ = e.NewContext(req, rec)

		claims, err := authHelper.VerifyJWT(req.Header)

		require.NotEmpty(t, claims)
		require.Empty(t, err)
	})

	t.Run("Should return if no auth", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		_ = e.NewContext(req, rec)

		claims, err := authHelper.VerifyJWT(req.Header)

		require.Empty(t, claims)
		require.NotEmpty(t, err)
	})

	t.Run("Should return error if auth not valid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+faker.Jwt())
		rec := httptest.NewRecorder()
		_ = e.NewContext(req, rec)

		claims, err := authHelper.VerifyJWT(req.Header)

		require.Empty(t, claims)
		require.NotEmpty(t, err)
	})

	t.Run("Should return error if token expired", func(t *testing.T) {
		expiredJWt := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDcwOTQ3MTQsImlhdCI6MTcwNzAwODMxNCwiaXNzIjoiYjJiNGYzMjItNWRmYS00NzcxLWE0MDYtODMyODBkZmFlNWYyIiwicm9sZXMiOlsiQURNSU5JU1RSQVRPUiJdLCJzdWIiOiJhZG1pbiIsInVzciI6IkFkbWluaXN0cmF0b3IifQ.ZUBJTFfJXqTXiv_WcQyjA4UBilJCHjIyyKz8uXsXq6U`
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+expiredJWt)
		rec := httptest.NewRecorder()
		_ = e.NewContext(req, rec)

		claims, err := authHelper.VerifyJWT(req.Header)

		require.Empty(t, claims)
		require.NotEmpty(t, err)
	})
}
