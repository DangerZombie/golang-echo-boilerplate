package repository

import "gorm.io/gorm"

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
