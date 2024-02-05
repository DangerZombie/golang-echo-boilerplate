package service_user_test

import (
	"errors"
	"go-echo/helper/auth"
	"go-echo/initialization"
	"go-echo/model/parameter"
	"go-echo/model/request"
	"go-echo/repository"
	"go-echo/repository/repository_user"
	"go-echo/service/service_user"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestLogin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)
	mockUserRepository := repository_user.NewMockUserRepository(mockCtrl)
	mockBaseRepository := repository.NewMockBaseRepository(mockCtrl)

	var logger *zap.Logger
	logger, _ = initialization.NewZapLogger("")
	userService := service_user.NewUserService(
		logger,
		mockAuthHelper,
		mockBaseRepository,
		mockUserRepository)

	username := faker.Name()
	password := faker.Password()
	token := faker.Jwt()

	loginRequest := request.LoginRequestBody{
		Username: username,
		Password: password,
	}

	loginEmptyRequest := request.LoginRequestBody{
		Username: "",
		Password: "",
	}

	findUserByUsernameAndPasswordInput := parameter.FindUserByUsernameAndPasswordInput{
		Username: username,
		Password: password,
	}

	findUserByUsernameAndPasswordOutput := parameter.FindUserByUsernameAndPasswordOutput{
		Id:       faker.UUIDHyphenated(),
		Username: username,
		Password: password,
	}

	t.Run("Should return token", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserByUsernameAndPassword(gomock.Any(), findUserByUsernameAndPasswordInput).
			Times(1).
			Return(findUserByUsernameAndPasswordOutput, nil)

		mockAuthHelper.EXPECT().
			GenerateJWT(username).
			Times(1).
			Return(token, nil)

		result, message, err := userService.Login(loginRequest)

		require.Equal(t, token, result.Token)
		require.NotEmpty(t, message)
		require.Empty(t, err)
	})

	t.Run("Should return error if username or password empty", func(t *testing.T) {
		result, message, err := userService.Login(loginEmptyRequest)

		require.Empty(t, result)
		require.NotEmpty(t, message)
		require.NotEmpty(t, err)
	})

	t.Run("Should return error if failed to get user", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserByUsernameAndPassword(gomock.Any(), findUserByUsernameAndPasswordInput).
			Times(1).
			Return(parameter.FindUserByUsernameAndPasswordOutput{}, errors.New("failed"))

		result, message, err := userService.Login(loginRequest)

		require.Empty(t, result)
		require.NotEmpty(t, message)
		require.NotEmpty(t, err)
	})

	t.Run("Should return error if failed to generate token", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserByUsernameAndPassword(gomock.Any(), findUserByUsernameAndPasswordInput).
			Times(1).
			Return(findUserByUsernameAndPasswordOutput, nil)

		mockAuthHelper.EXPECT().
			GenerateJWT(username).
			Times(1).
			Return("", errors.New("failed"))

		result, message, err := userService.Login(loginRequest)

		require.Empty(t, result)
		require.NotEmpty(t, message)
		require.NotEmpty(t, err)
	})
}
