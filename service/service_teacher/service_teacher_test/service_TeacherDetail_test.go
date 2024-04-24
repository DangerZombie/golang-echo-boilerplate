package service_teacher_test

import (
	"errors"
	"go-echo/helper/message"
	"go-echo/initialization"
	"go-echo/model/parameter"
	"go-echo/model/request"
	"go-echo/repository"
	"go-echo/repository/repository_teacher"
	"go-echo/service/service_teacher"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestTeacherDetail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockBaseRepository := repository.NewMockBaseRepository(mockCtrl)
	mockTeacherRepository := repository_teacher.NewMockTeacherRepository(mockCtrl)

	var logger *zap.Logger
	logger, _ = initialization.NewZapLogger("")
	teacherService := service_teacher.NewTeacherService(
		logger,
		mockBaseRepository,
		mockTeacherRepository)

	id := faker.UUIDHyphenated()
	nickname := faker.Name()
	email := faker.Email()
	jobTitleId := faker.UUIDHyphenated()
	teacherDetailRequest := request.TeacherDetailRequest{
		Id: id,
	}

	teacherDetailEmptyRequest := request.TeacherDetailRequest{
		Id: "",
	}

	findTeacherByIdInput := parameter.FindTeacherByIdInput{
		Id: id,
	}

	findTeacherByIdOutput := parameter.FindTeacherByIdOutput{
		Id:           id,
		Nickname:     nickname,
		Email:        email,
		Status:       "PERMANENT",
		Experience:   10,
		Degree:       "S.Pd",
		JobTitleId:   jobTitleId,
		JobTitleName: "Principal",
	}

	t.Run("Should return data teacher", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockTeacherRepository.EXPECT().
			FindTeacherById(gomock.Any(), findTeacherByIdInput).
			Times(1).
			Return(findTeacherByIdOutput, nil)

		mockBaseRepository.EXPECT().
			BeginCommit(gomock.Any()).
			Times(1).
			Return()

		result, msg, errMsg := teacherService.TeacherDetail(teacherDetailRequest)

		require.Equal(t, id, result.Id)
		require.Equal(t, email, result.Email)
		require.Equal(t, nickname, result.Nickname)
		require.Equal(t, jobTitleId, result.JobTitleId)
		require.Equal(t, "PERMANENT", result.Status)
		require.Equal(t, "S.Pd", result.Degree)
		require.Equal(t, 10, result.Experience)
		require.Equal(t, "Principal", result.JobTitleName)
		require.Equal(t, message.SuccessMsg, msg)
		require.Empty(t, errMsg)
	})

	t.Run("Should return error if id is empty", func(t *testing.T) {
		result, msg, errMsg := teacherService.TeacherDetail(teacherDetailEmptyRequest)

		require.Empty(t, result)
		require.Equal(t, message.FailedMsg, msg)
		require.NotEmpty(t, errMsg)
	})

	t.Run("Should return error if failed to get data teacher", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockTeacherRepository.EXPECT().
			FindTeacherById(gomock.Any(), findTeacherByIdInput).
			Times(1).
			Return(parameter.FindTeacherByIdOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, msg, errMsg := teacherService.TeacherDetail(teacherDetailRequest)

		require.Empty(t, result)
		require.Equal(t, message.FailedMsg, msg)
		require.NotEmpty(t, errMsg)
	})
}
