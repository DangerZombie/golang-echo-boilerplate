package base

import (
	"time"

	"gorm.io/gorm"

	gouuid "github.com/google/uuid"
)

type BaseModel struct {
	// ID
	// in: string
	ID string `gorm:"primary_key" json:"id"`

	// Created At
	// in: int64
	CreatedAt int64 `gorm:"type:bigint" json:"created_at"`

	// Created By
	// in: int64
	CreatedBy string `gorm:"type:varchar(50)" json:"created_by"`

	// Updated At
	// in: int64
	UpdatedAt int64 `gorm:"type:bigint" json:"updated_at"`

	// Updated By
	// in: string
	UpdatedBy string `gorm:"type:varchar(50)" json:"updated_by"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	uuid := gouuid.New()
	time := time.Now().Unix()
	tx.Statement.SetColumn("ID", uuid)
	tx.Statement.SetColumn("CreatedAt", time)
	tx.Statement.SetColumn("UpdatedAt", time)
	return nil
}

func (base *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	time := time.Now().Unix()
	tx.Statement.SetColumn("UpdatedAt", time)
	return nil
}
