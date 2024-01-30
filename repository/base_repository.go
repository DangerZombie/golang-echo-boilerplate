package repository

import (
	"gorm.io/gorm"
)

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
