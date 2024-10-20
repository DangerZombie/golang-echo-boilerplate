package repository_teacher

import "go-echo/repository"

type teacherRepo struct {
	base repository.BaseRepository
}

func NewTeacherRepository(br repository.BaseRepository) TeacherRepository {
	return &teacherRepo{br}
}
