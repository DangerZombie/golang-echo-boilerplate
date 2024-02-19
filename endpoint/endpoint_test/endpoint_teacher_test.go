package endpoint_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-echo/endpoint"
	"go-echo/helper/auth"
	"go-echo/helper/message"
	"go-echo/helper/static"
	"go-echo/model/base"
	"go-echo/model/parameter"
	"go-echo/model/request"
	"go-echo/model/response"
	"go-echo/service/service_teacher"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestEndpoint_CreateTeacher(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockTeacherService := service_teacher.NewMockTeacherService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	e := echo.New()
	id := faker.UUIDHyphenated()
	userId := faker.UUIDHyphenated()
	jobTitleId := faker.UUIDHyphenated()
	status := "PERMANENT"
	experience := 10
	degree := "S.Pd"
	issuerId := faker.UUIDHyphenated()
	username := faker.Email()
	nickname := faker.Name()
	claims := parameter.JwtClaims{
		Issuer:  issuerId,
		Subject: username,
		User:    nickname,
		Roles:   []string{static.RoleADMINISTRATOR},
	}

	teacherCreateRequest := request.TeacherCreateRequestBody{
		UserId:     userId,
		JobTitleId: jobTitleId,
		Status:     status,
		Experience: experience,
		Degree:     degree,
		Issuer:     issuerId,
	}

	teacherCreateResponse := response.TeacherCreateResponse{
		Id: id,
	}

	t.Run("Should return OK", func(t *testing.T) {
		reqBody, _ := json.Marshal(teacherCreateRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/teacher", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockTeacherService.EXPECT().
			TeacherCreate(teacherCreateRequest).
			Times(1).
			Return(teacherCreateResponse, message.SuccessMsg, nil)

		statusCode, result := endpointModule.CreateTeacherRequest(c, mockTeacherService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Unauthorized", func(t *testing.T) {
		reqBody, _ := json.Marshal(teacherCreateRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/teacher", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(parameter.JwtClaims{}, errors.New("failed"))

		statusCode, result := endpointModule.CreateTeacherRequest(c, mockTeacherService)

		require.Equal(t, http.StatusUnauthorized, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Forbidden", func(t *testing.T) {
		reqBody, _ := json.Marshal(teacherCreateRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/teacher", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(parameter.JwtClaims{}, nil)

		statusCode, result := endpointModule.CreateTeacherRequest(c, mockTeacherService)

		require.Equal(t, http.StatusForbidden, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Bad Request", func(t *testing.T) {
		reqBody, _ := json.Marshal(teacherCreateRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/teacher", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockTeacherService.EXPECT().
			TeacherCreate(teacherCreateRequest).
			Times(1).
			Return(response.TeacherCreateResponse{}, message.ErrReqParam, nil)

		statusCode, result := endpointModule.CreateTeacherRequest(c, mockTeacherService)

		require.Equal(t, http.StatusBadRequest, statusCode)
		require.NotEmpty(t, result)
	})
}

func TestEndpoint_ListTeachers(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockTeacherService := service_teacher.NewMockTeacherService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	e := echo.New()
	issuerId := faker.UUIDHyphenated()
	username := faker.Email()
	nickname := faker.Name()
	claims := parameter.JwtClaims{
		Issuer:  issuerId,
		Subject: username,
		User:    nickname,
		Roles:   []string{static.RoleADMINISTRATOR},
	}

	teacherListRequest := request.TeacherListRequest{
		Page:  1,
		Limit: 10,
		Sort:  "created_at_utc0",
		Dir:   "asc",
		Name:  "John",
	}

	teacherListRespose := []response.TeacherListResponse{
		{
			Id:         faker.UUIDHyphenated(),
			Nickname:   "John",
			Email:      faker.Email(),
			Status:     "PERMANENT",
			Experience: 10,
			Degree:     "S.Pd",
		},
	}

	paginationResponse := base.Pagination{
		Records:      0,
		TotalRecords: 1,
		Limit:        10,
		Page:         1,
		TotalPage:    1,
	}

	t.Run("Should return OK", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/teacher?page=1&limit=10&sort=created_at_utc0&dir=asc&name=John", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockTeacherService.EXPECT().
			TeacherList(teacherListRequest).
			Times(1).
			Return(teacherListRespose, paginationResponse, message.SuccessMsg, nil)

		statusCode, result := endpointModule.ListTeachersRequest(c, mockTeacherService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Unauthorized", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/teacher?page=1&limit=10&sort=created_at_utc0&dir=asc&name=John", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(parameter.JwtClaims{}, errors.New("failed"))

		statusCode, result := endpointModule.ListTeachersRequest(c, mockTeacherService)

		require.Equal(t, http.StatusUnauthorized, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Bad Request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/teacher?page=1&limit=10&sort=created_at_utc0&dir=asc&name=John", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockTeacherService.EXPECT().
			TeacherList(teacherListRequest).
			Times(1).
			Return([]response.TeacherListResponse{}, base.Pagination{}, message.ErrReqParam, nil)

		statusCode, result := endpointModule.ListTeachersRequest(c, mockTeacherService)

		require.Equal(t, http.StatusBadRequest, statusCode)
		require.NotEmpty(t, result)
	})
}

func TestEndpoint_FindTeacherDetail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockTeacherService := service_teacher.NewMockTeacherService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	e := echo.New()
	issuerId := faker.UUIDHyphenated()
	username := faker.Email()
	nickname := faker.Name()
	id := faker.UUIDHyphenated()
	claims := parameter.JwtClaims{
		Issuer:  issuerId,
		Subject: username,
		User:    nickname,
		Roles:   []string{static.RoleADMINISTRATOR},
	}

	teacherDetailRequest := request.TeacherDetailRequest{
		Id: id,
	}

	teacherDetailResponse := response.TeacherDetailResponse{
		Id:           id,
		Nickname:     faker.Name(),
		Email:        faker.Email(),
		Status:       "PERMANENT",
		Experience:   10,
		Degree:       "S.Pd",
		JobTitleId:   faker.UUIDHyphenated(),
		JobTitleName: "Principal",
	}

	t.Run("Should return OK", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/teacher/%s", id)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/teacher/:id")
		c.SetParamNames("id")
		c.SetParamValues(id)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockTeacherService.EXPECT().
			TeacherDetail(teacherDetailRequest).
			Times(1).
			Return(teacherDetailResponse, message.SuccessMsg, nil)

		statusCode, result := endpointModule.FindTeacherDetailRequest(c, mockTeacherService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Unauthorized", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/teacher/%s", id)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(parameter.JwtClaims{}, errors.New("failed"))

		statusCode, result := endpointModule.FindTeacherDetailRequest(c, mockTeacherService)

		require.Equal(t, http.StatusUnauthorized, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Bad Request", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/teacher/%s", id)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockTeacherService.EXPECT().
			TeacherDetail(gomock.Any()).
			Times(1).
			Return(response.TeacherDetailResponse{}, message.ErrReqParam, nil)

		statusCode, result := endpointModule.FindTeacherDetailRequest(c, mockTeacherService)

		require.Equal(t, http.StatusBadRequest, statusCode)
		require.NotEmpty(t, result)
	})
}

func TestEndpoint_UpdateTeacher(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockTeacherService := service_teacher.NewMockTeacherService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	e := echo.New()
	issuerId := faker.UUIDHyphenated()
	username := faker.Email()
	nickname := faker.Name()
	id := faker.UUIDHyphenated()
	jobTitleId := faker.UUIDHyphenated()
	status := "PERMANENT"
	experience := 10
	degree := "S.Pd"
	claims := parameter.JwtClaims{
		Issuer:  issuerId,
		Subject: username,
		User:    nickname,
		Roles:   []string{static.RoleADMINISTRATOR},
	}

	teacherUpdateBodyRequest := request.TeacherUpdateRequestBody{
		JobTitleId: &jobTitleId,
		Status:     &status,
		Experience: &experience,
		Degree:     &degree,
	}

	teacherUpdateRequest := request.TeacherUpdateRequest{
		Id:   id,
		Body: teacherUpdateBodyRequest,
	}

	teacherUpdateResponse := response.TeacherUpdateResponse{
		Id:         id,
		JobTitleId: faker.UUIDHyphenated(),
		Status:     "PERMANENT",
		Experience: 10,
		Degree:     "S.Pd",
	}

	t.Run("Should return OK", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/teacher/%s", id)
		reqBody, _ := json.Marshal(teacherUpdateBodyRequest)
		req := httptest.NewRequest(http.MethodPut, url, strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/teacher/:id")
		c.SetParamNames("id")
		c.SetParamValues(id)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockTeacherService.EXPECT().
			TeacherUpdate(teacherUpdateRequest).
			Times(1).
			Return(teacherUpdateResponse, message.SuccessMsg, nil)

		statusCode, result := endpointModule.UpdateTeacherRequest(c, mockTeacherService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Unauthorized", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/teacher/%s", id)
		reqBody, _ := json.Marshal(teacherUpdateBodyRequest)
		req := httptest.NewRequest(http.MethodPut, url, strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(parameter.JwtClaims{}, errors.New("failed"))

		statusCode, result := endpointModule.UpdateTeacherRequest(c, mockTeacherService)

		require.Equal(t, http.StatusUnauthorized, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Bad Request", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/teacher/%s", id)
		reqBody, _ := json.Marshal(teacherUpdateBodyRequest)
		req := httptest.NewRequest(http.MethodPut, url, strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockTeacherService.EXPECT().
			TeacherUpdate(gomock.Any()).
			Times(1).
			Return(response.TeacherUpdateResponse{}, message.ErrReqParam, nil)

		statusCode, result := endpointModule.UpdateTeacherRequest(c, mockTeacherService)

		require.Equal(t, http.StatusBadRequest, statusCode)
		require.NotEmpty(t, result)
	})
}

func TestEndpoint_DeleteTeacher(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockTeacherService := service_teacher.NewMockTeacherService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	e := echo.New()
	issuerId := faker.UUIDHyphenated()
	username := faker.Email()
	nickname := faker.Name()
	id := faker.UUIDHyphenated()
	claims := parameter.JwtClaims{
		Issuer:  issuerId,
		Subject: username,
		User:    nickname,
		Roles:   []string{static.RoleADMINISTRATOR},
	}

	teacherDeleteRequest := request.TeacherDeleteRequest{
		Id: id,
	}

	teacherDeleteResponse := response.TeacherDeleteResponse{
		Success: true,
	}

	t.Run("Should return OK", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/teacher/%s", id)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/teacher/:id")
		c.SetParamNames("id")
		c.SetParamValues(id)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockTeacherService.EXPECT().
			TeacherDelete(teacherDeleteRequest).
			Times(1).
			Return(teacherDeleteResponse, message.SuccessMsg, nil)

		statusCode, result := endpointModule.DeleteTeacherRequest(c, mockTeacherService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Unauthorized", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/teacher/%s", id)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(parameter.JwtClaims{}, errors.New("failed"))

		statusCode, result := endpointModule.DeleteTeacherRequest(c, mockTeacherService)

		require.Equal(t, http.StatusUnauthorized, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Bad Request", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/teacher/%s", id)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		req.Header.Set(echo.HeaderAuthorization, faker.JWT)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockTeacherService.EXPECT().
			TeacherDelete(gomock.Any()).
			Times(1).
			Return(response.TeacherDeleteResponse{}, message.ErrReqParam, nil)

		statusCode, result := endpointModule.DeleteTeacherRequest(c, mockTeacherService)

		require.Equal(t, http.StatusBadRequest, statusCode)
		require.NotEmpty(t, result)
	})
}
