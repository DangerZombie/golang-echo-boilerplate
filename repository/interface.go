package repository

import "gorm.io/gorm"

type BaseRepository interface {
	GetDB() *gorm.DB
	GetBegin() *gorm.DB
	BeginCommit(db *gorm.DB)
	BeginRollback(db *gorm.DB)
}
