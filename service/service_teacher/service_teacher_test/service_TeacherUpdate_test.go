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

func TestTeacherUpdate(t *testing.T) {
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
	jobTitleId := faker.UUIDHyphenated()
	status := "PERMANTENT"
	experience := 10
	degree := "S.Pd"
	updateTeacherRequest := request.TeacherUpdateRequest{
		Id: id,
		Body: request.TeacherUpdateRequestBody{
			JobTitleId: &jobTitleId,
			Status:     &status,
			Experience: &experience,
			Degree:     &degree,
		},
	}

	updateTeacherEmptyRequest := request.TeacherUpdateRequest{
		Id: "",
	}

	updateTeacherByIdInput := parameter.UpdateTeacherByIdInput{
		Id: id,
		Fields: map[string]interface{}{
			"job_title_id": &updateTeacherRequest.Body.JobTitleId,
			"status":       &updateTeacherRequest.Body.Status,
			"experience":   &updateTeacherRequest.Body.Experience,
			"degree":       &updateTeacherRequest.Body.Degree,
		},
	}

	updateTeacherByIdOutput := parameter.UpdateTeacherByIdOutput{
		Teacher: entity.Teacher{
			BaseModel: base.BaseModel{
				Id: id,
			},
			Status:     status,
			JobTitleId: jobTitleId,
			Experience: experience,
			Degree:     degree,
		},
	}

	t.Run("Should return updated data", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockTeacherRepository.EXPECT().
			UpdateTeacherById(gomock.Any(), updateTeacherByIdInput).
			Times(1).
			Return(updateTeacherByIdOutput, nil)

		mockBaseRepository.EXPECT().
			BeginCommit(gomock.Any()).
			Times(1).
			Return()

		result, msg, errMsg := teacherService.TeacherUpdate(updateTeacherRequest)

		require.Equal(t, id, result.Id)
		require.Equal(t, jobTitleId, result.JobTitleId)
		require.Equal(t, status, result.Status)
		require.Equal(t, experience, result.Experience)
		require.Equal(t, degree, result.Degree)
		require.Equal(t, message.SuccessMsg, msg)
		require.Empty(t, errMsg)
	})

	t.Run("Should return error if id is empty", func(t *testing.T) {
		result, msg, errMsg := teacherService.TeacherUpdate(updateTeacherEmptyRequest)

		require.Empty(t, result)
		require.Equal(t, message.FailedMsg, msg)
		require.NotEmpty(t, errMsg)
	})

	t.Run("Should return error if failed to update teacher", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockTeacherRepository.EXPECT().
			UpdateTeacherById(gomock.Any(), updateTeacherByIdInput).
			Times(1).
			Return(parameter.UpdateTeacherByIdOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, msg, errMsg := teacherService.TeacherUpdate(updateTeacherRequest)

		require.Empty(t, result)
		require.Equal(t, message.FailedMsg, msg)
		require.NotEmpty(t, errMsg)
	})
}
