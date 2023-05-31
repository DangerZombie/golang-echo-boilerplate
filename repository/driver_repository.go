package repository

import (
	"errors"
	"fmt"
	"go-echo/helper/util"
	"go-echo/model/base"
	"go-echo/model/entity"
	"math"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type driverRepo struct {
	base BaseRepository
}

type DriverRepository interface {
	InsertDriver(db *gorm.DB, driver entity.Driver) (*entity.Driver, error)
	GetListDrivers(db *gorm.DB, limit int, page int, sort string, dir string, filter map[string]interface{}) ([]entity.Driver, *base.Pagination, error)
	GetDriverByNumber(db *gorm.DB, licenseNumber string) (*entity.Driver, error)
	UpdateDriverByNumber(db *gorm.DB, licenseNumber string, input map[string]interface{}) (*entity.Driver, error)
	DeleteDriverByNumber(db *gorm.DB, licenseNumber string) error
}

func NewDriverRepository(br BaseRepository) DriverRepository {
	return &driverRepo{br}
}

func (r *driverRepo) InsertDriver(db *gorm.DB, driver entity.Driver) (*entity.Driver, error) {
	err := db.Create(&driver).Error
	if err != nil {
		return nil, err
	}

	return &driver, nil
}

func (r *driverRepo) GetListDrivers(db *gorm.DB, limit int, page int, sort string, dir string, filter map[string]interface{}) ([]entity.Driver, *base.Pagination, error) {
	var drivers []entity.Driver
	var pagination base.Pagination

	query := db

	if filter["name"] != "" {
		query = query.Where("name = ?", filter["name"].(string))
	}

	pagination.Limit = limit
	pagination.Page = page

	sort = util.ReplaceEmptyString(sort, "created_at")
	dir = util.ReplaceEmptyString(dir, "desc")
	if strings.EqualFold(dir, "asc") {
		dir = "asc"
	}

	query = query.Order(fmt.Sprintf("%s %s", sort, dir))

	err := query.Scopes(r.Paginate(drivers, &pagination, query, int64(len(drivers)))).
		Find(&drivers).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	return drivers, &pagination, nil
}

func (r *driverRepo) GetDriverByNumber(db *gorm.DB, licenseNumber string) (*entity.Driver, error) {
	var driver entity.Driver
	err := db.
		Model(&entity.Driver{}).
		Where("license_number = ?", licenseNumber).
		First(&driver).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &driver, nil
}

func (r *driverRepo) UpdateDriverByNumber(db *gorm.DB, licenseNumber string, input map[string]interface{}) (*entity.Driver, error) {
	var driver entity.Driver
	err := db.
		Model(&driver).
		Where("license_number = ?", licenseNumber).
		Clauses(clause.Returning{}).
		Updates(input).
		Error

	if err != nil {
		return nil, err
	}

	return &driver, nil
}

func (r *driverRepo) DeleteDriverByNumber(db *gorm.DB, licenseNumber string) error {
	err := db.
		Where("license_number = ?", licenseNumber).
		Delete(&entity.Driver{}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *driverRepo) Paginate(value interface{}, pagination *base.Pagination, db *gorm.DB, currRecord int64) func(db *gorm.DB) *gorm.DB {
	var totalRecords int64
	db.Model(value).Count(&totalRecords)

	pagination.TotalRecords = totalRecords
	pagination.TotalPage = int(math.Ceil(float64(totalRecords) / float64(pagination.GetLimit())))
	pagination.Records = int64(pagination.Limit*(pagination.Page-1)) + int64(currRecord)

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
