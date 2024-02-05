package endpoint_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-echo/endpoint"
	"go-echo/helper/auth"
	"go-echo/helper/message"
	"go-echo/helper/static"
	"go-echo/model/parameter"
	"go-echo/model/request"
	"go-echo/model/response"
	"go-echo/service/service_user"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestEndpointUser_Login(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockUserService := service_user.NewMockUserService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	e := echo.New()
	loginRequest := request.LoginRequestBody{
		Username: faker.Username(),
		Password: faker.Name(),
	}

	loginResponse := response.LoginResponse{
		Token: faker.Jwt(),
	}

	t.Run("Should return OK", func(t *testing.T) {
		reqBody, _ := json.Marshal(loginRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/login", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockUserService.EXPECT().
			Login(loginRequest).
			Times(1).
			Return(loginResponse, message.SuccessMsg, nil)

		statusCode, result := endpointModule.LoginRequest(c, mockUserService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Bad Request if something went wrong", func(t *testing.T) {
		reqBody, _ := json.Marshal(loginRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/login", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockUserService.EXPECT().
			Login(loginRequest).
			Times(1).
			Return(response.LoginResponse{}, message.ErrReqParam, nil)

		statusCode, result := endpointModule.LoginRequest(c, mockUserService)

		require.Equal(t, http.StatusBadRequest, statusCode)
		require.NotEmpty(t, result)
	})
}

func TestEndpointUser_UserProfile(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockUserService := service_user.NewMockUserService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	e := echo.New()
	id := faker.UUIDHyphenated()

	userProfileInput := request.UserProfileRequest{
		Id: id,
	}

	userProfileOutput := response.UserProfileResponse{
		Id:       id,
		Nickname: faker.Name(),
	}

	t.Run("Should return OK", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/user/profile?id=%s", id)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(parameter.JwtClaims{}, nil)

		mockUserService.EXPECT().
			UserProfile(userProfileInput).
			Times(1).
			Return(userProfileOutput, message.SuccessMsg, nil)

		statusCode, result := endpointModule.UserProfileRequest(c, mockUserService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Unauthorized", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/user/profile?id=%s", id)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(parameter.JwtClaims{}, errors.New("failed"))

		statusCode, result := endpointModule.UserProfileRequest(c, mockUserService)

		require.Equal(t, http.StatusUnauthorized, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Bad Request", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/user/profile?id=%s", id)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(parameter.JwtClaims{}, nil)

		mockUserService.EXPECT().
			UserProfile(userProfileInput).
			Times(1).
			Return(response.UserProfileResponse{}, message.ErrReqParam, nil)

		statusCode, result := endpointModule.UserProfileRequest(c, mockUserService)

		require.Equal(t, http.StatusBadRequest, statusCode)
		require.NotEmpty(t, result)
	})
}

func TestEndpointUser_RegisterUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockUserService := service_user.NewMockUserService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	e := echo.New()
	id := faker.UUIDHyphenated()
	username := faker.Username()
	password := faker.Password()
	nickname := faker.Name()

	claims := parameter.JwtClaims{
		Issuer:  faker.UUIDHyphenated(),
		Subject: username,
		User:    nickname,
		Roles:   []string{static.RoleADMINISTRATOR},
	}

	claimsWrongRole := parameter.JwtClaims{
		Issuer:  faker.UUIDHyphenated(),
		Subject: username,
		User:    nickname,
		Roles:   []string{faker.Username()},
	}

	registerUserInput := request.RegisterUserRequestBody{
		Username: username,
		Password: password,
		Nickname: nickname,
	}

	registerUserOutput := response.RegisterUserResponse{
		Id: id,
	}

	t.Run("Should return OK", func(t *testing.T) {
		reqBody, _ := json.Marshal(registerUserInput)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/register", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockUserService.EXPECT().
			RegisterUser(registerUserInput).
			Times(1).
			Return(registerUserOutput, message.SuccessMsg, nil)

		statusCode, result := endpointModule.RegisterUserRequest(c, mockUserService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Unauthorized", func(t *testing.T) {
		reqBody, _ := json.Marshal(registerUserInput)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/register", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(parameter.JwtClaims{}, errors.New("failed"))

		statusCode, result := endpointModule.RegisterUserRequest(c, mockUserService)

		require.Equal(t, http.StatusUnauthorized, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Forbidden", func(t *testing.T) {
		reqBody, _ := json.Marshal(registerUserInput)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/register", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claimsWrongRole, nil)

		statusCode, result := endpointModule.RegisterUserRequest(c, mockUserService)

		require.Equal(t, http.StatusForbidden, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Bad Request", func(t *testing.T) {
		reqBody, _ := json.Marshal(registerUserInput)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/register", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockUserService.EXPECT().
			RegisterUser(registerUserInput).
			Times(1).
			Return(response.RegisterUserResponse{}, message.ErrReqParam, nil)

		statusCode, result := endpointModule.RegisterUserRequest(c, mockUserService)

		require.Equal(t, http.StatusBadRequest, statusCode)
		require.NotEmpty(t, result)
	})
}
