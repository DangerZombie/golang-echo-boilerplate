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

func TestTeacherList(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockBaseRepository := repository.NewMockBaseRepository(mockCtrl)
	mockTeacherRepository := repository_teacher.NewMockTeacherRepository(mockCtrl)

	var logger *zap.Logger
	logger, _ = initialization.NewZapLogger("")
	teacherService := service_teacher.NewTeacherService(
		logger,
		mockBaseRepository,
		mockTeacherRepository)

	page, limit := 0, 0
	sort, dir := "create_at_utc0", "asc"
	name := faker.Name()
	teacherListRequest := request.TeacherListRequest{
		Page:  page,
		Limit: limit,
		Sort:  sort,
		Dir:   dir,
		Name:  name,
	}

	filter := map[string]interface{}{
		"name": name,
	}

	teacherListInput := parameter.ListTeacherInput{
		Limit:  10,
		Page:   1,
		Sort:   sort,
		Dir:    dir,
		Filter: filter,
	}

	teacherListOutput := parameter.ListTeacherOutput{
		Items: []entity.Teacher{
			{
				BaseModel: base.BaseModel{
					Id: faker.UUIDHyphenated(),
				},
				UserId:     faker.UUIDDigit(),
				Status:     "PERMANENT",
				Experience: 10,
				Degree:     "S.Pd",
				JobTitleId: faker.UUIDHyphenated(),
				JobTitle: entity.JobTitle{
					Name: "Principal",
				},
			},
		},
	}

	paginationOutput := base.Pagination{
		Records:      0,
		TotalRecords: 1,
		Limit:        10,
		Page:         1,
		TotalPage:    1,
	}

	t.Run("Should return list teacher", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockTeacherRepository.EXPECT().
			ListTeacher(gomock.Any(), teacherListInput).
			Times(1).
			Return(teacherListOutput, paginationOutput, nil)

		mockBaseRepository.EXPECT().
			BeginCommit(gomock.Any()).
			Times(1).
			Return()

		result, pagination, msg, errMsg := teacherService.TeacherList(teacherListRequest)

		require.Equal(t, 1, len(result))
		require.Equal(t, paginationOutput, pagination)
		require.Equal(t, message.SuccessMsg, msg)
		require.Empty(t, errMsg)
	})

	t.Run("Should return error if failed to get list teacher", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockTeacherRepository.EXPECT().
			ListTeacher(gomock.Any(), teacherListInput).
			Times(1).
			Return(parameter.ListTeacherOutput{}, base.Pagination{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, pagination, msg, errMsg := teacherService.TeacherList(teacherListRequest)

		require.Empty(t, result)
		require.Empty(t, pagination)
		require.Equal(t, message.FailedMsg, msg)
		require.NotEmpty(t, errMsg)
	})
}
