package auth_test

import (
	"errors"
	"go-echo/helper/auth"
	"go-echo/helper/static"
	"go-echo/model/base"
	"go-echo/model/entity"
	"go-echo/model/parameter"
	"go-echo/repository"
	"go-echo/repository/repository_user"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestAuthGenerateJWT(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockBaseRepository := repository.NewMockBaseRepository(mockCtrl)
	mockUserRepository := repository_user.NewMockUserRepository(mockCtrl)

	authHelper := auth.NewAuthHelper(
		mockBaseRepository,
		mockUserRepository,
	)

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

	t.Run("Should return token", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserRoleByUserId(gomock.Any(), findUserRoleByUserIdInput).
			Times(1).
			Return(findUserRoleByUserIdOutput, nil)

		result, err := authHelper.GenerateJWT(id)

		require.NotEmpty(t, result)
		require.Empty(t, err)
	})

	t.Run("Should return error if failed to fetch user", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserRoleByUserId(gomock.Any(), findUserRoleByUserIdInput).
			Times(1).
			Return(parameter.FindUserRoleByUserIdOutput{}, errors.New("failed"))

		result, err := authHelper.GenerateJWT(id)

		require.Empty(t, result)
		require.NotEmpty(t, err)
	})
}
