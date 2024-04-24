package repository

import "gorm.io/gorm"

type baseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &baseRepository{db}
}
