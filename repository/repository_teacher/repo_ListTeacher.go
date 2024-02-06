package repository_teacher

import (
	"errors"
	"fmt"
	"go-echo/helper/util"
	"go-echo/model/base"
	"go-echo/model/entity"
	"go-echo/model/parameter"
	"math"
	"strings"

	"gorm.io/gorm"
)

func (r *teacherRepo) ListTeacher(db *gorm.DB, input parameter.ListTeacherInput) (output parameter.ListTeacherOutput, pagination base.Pagination, err error) {
	var teachers []entity.Teacher

	query := db.Debug()

	if input.Filter["name"] != "" {
		query = query.Where("user.username = ?", input.Filter["name"].(string))
	}

	pagination.Limit = input.Limit
	pagination.Page = input.Page

	input.Sort = util.ReplaceEmptyString(input.Sort, "teacher.created_at_utc0")
	input.Dir = util.ReplaceEmptyString(input.Dir, "desc")
	if strings.EqualFold(input.Dir, "asc") {
		input.Dir = "asc"
	}

	query = query.Order(fmt.Sprintf("%s %s", input.Sort, input.Dir))

	err = query.
		Scopes(r.Paginate(teachers, &pagination, query, int64(len(teachers)))).
		Preload("User").
		Find(&teachers).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}

		return
	}

	output.Items = teachers

	return output, pagination, nil
}

func (r *teacherRepo) Paginate(value interface{}, pagination *base.Pagination, db *gorm.DB, currRecord int64) func(db *gorm.DB) *gorm.DB {
	var totalRecords int64
	db.Model(value).Count(&totalRecords)

	pagination.TotalRecords = totalRecords
	pagination.TotalPage = int(math.Ceil(float64(totalRecords) / float64(pagination.GetLimit())))
	pagination.Records = int64(pagination.Limit*(pagination.Page-1)) + int64(currRecord)

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
