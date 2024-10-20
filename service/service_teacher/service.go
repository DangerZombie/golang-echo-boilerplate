package service_teacher

import (
	"go-echo/repository"
	"go-echo/repository/repository_teacher"

	"go.uber.org/zap"
)

type teacherServiceImpl struct {
	logger      *zap.Logger
	baseRepo    repository.BaseRepository
	teacherRepo repository_teacher.TeacherRepository
}

func NewTeacherService(
	lg *zap.Logger,
	br repository.BaseRepository,
	tr repository_teacher.TeacherRepository,
) TeacherService {
	return &teacherServiceImpl{lg, br, tr}
}
