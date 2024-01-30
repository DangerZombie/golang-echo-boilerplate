package service_user_test

import (
	"go-echo/helper/auth"
	"go-echo/initialization"
	"go-echo/model/entity"
	"go-echo/model/request"
	"go-echo/repository"
	"go-echo/repository/repository_user"
	"go-echo/service/service_user"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"gorm.io/gorm"
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

	loginRequest := request.LoginRequest{
		Username: username,
		Password: password,
	}

	findUserByUsernameAndPasswordOutput := entity.User{
		Id:       faker.UUIDHyphenated(),
		Username: username,
		Password: password,
	}

	t.Run("Should return token", func(t *testing.T) {
		var tx *gorm.DB
		mockUserRepository.EXPECT().
			FindUserByUsernameAndPassword(gomock.Any(), username, password).
			Times(1).
			Return(&findUserByUsernameAndPasswordOutput, nil)

		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(tx)

		mockAuthHelper.EXPECT().
			GenerateJWT(username).
			Times(1).
			Return(token, nil)

		result, message, err := userService.Login(loginRequest)

		require.Equal(t, token, result.Token)
		require.NotEmpty(t, message)
		require.Empty(t, err)
	})
}
