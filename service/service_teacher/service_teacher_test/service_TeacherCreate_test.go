package service_teacher_test

import (
	"errors"
	"go-echo/helper/message"
	"go-echo/initialization"
	"go-echo/model/base"
	"go-echo/model/entity"
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

func TestTeacherCreate(t *testing.T) {
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
	userId := faker.UUIDHyphenated()
	jobTitleId := faker.UUIDHyphenated()
	issuerId := faker.UUIDHyphenated()

	teacherCreateRequest := request.TeacherCreateRequestBody{
		UserId:     userId,
		JobTitleId: jobTitleId,
		Status:     "PERMANENT",
		Experience: 10,
		Degree:     "S.Pd",
		Issuer:     issuerId,
	}

	teacherInput := parameter.CreateTeacherInput{
		Teacher: entity.Teacher{
			UserId:     userId,
			JobTitleId: jobTitleId,
			Status:     "PERMANENT",
			Experience: 10,
			Degree:     "S.Pd",
			BaseModel: base.BaseModel{
				CreatedBy: issuerId,
				UpdatedBy: issuerId,
			},
		},
	}

	teacherOutput := parameter.CreateTeacherOutput{
		Id: id,
	}

	t.Run("Should return id if success", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockTeacherRepository.EXPECT().
			CreateTeacher(gomock.Any(), teacherInput).
			Times(1).
			Return(teacherOutput, nil)

		mockBaseRepository.EXPECT().
			BeginCommit(gomock.Any()).
			Times(1).
			Return()

		result, msg, errMsg := teacherService.TeacherCreate(teacherCreateRequest)

		require.Equal(t, id, result.Id)
		require.Equal(t, message.SuccessMsg, msg)
		require.Empty(t, errMsg)
	})

	t.Run("Should return error if failed to create teacher", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockTeacherRepository.EXPECT().
			CreateTeacher(gomock.Any(), teacherInput).
			Times(1).
			Return(parameter.CreateTeacherOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, msg, errMsg := teacherService.TeacherCreate(teacherCreateRequest)

		require.Empty(t, result.Id)
		require.Equal(t, message.FailedMsg, msg)
		require.NotEmpty(t, errMsg)
	})
}
