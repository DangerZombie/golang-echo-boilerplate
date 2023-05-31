package repository

import (
	"gorm.io/gorm"
)

type baseRepository struct {
	db *gorm.DB
}

type BaseRepository interface {
	GetDB() *gorm.DB
	GetBegin() *gorm.DB
	BeginCommit(db *gorm.DB)
	BeginRollback(db *gorm.DB)
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &baseRepository{db}
}

func (br *baseRepository) GetDB() *gorm.DB {
	return br.db
}

func (br baseRepository) GetBegin() *gorm.DB {
	return br.GetDB().Begin()
}

func (br baseRepository) BeginCommit(db *gorm.DB) {
	db.Commit()
}

func (br baseRepository) BeginRollback(db *gorm.DB) {
	db.Rollback()
}
