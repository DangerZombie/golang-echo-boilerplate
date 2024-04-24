package base

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	// ID
	// in: string
	Id string `gorm:"type:uuid;primaryKey" json:"id"`

	// Created At UTC0
	// in: int64
	CreatedAtUtc0 int64 `gorm:"type:int8" json:"created_at_utc0"`

	// Created By
	// in: string
	CreatedBy string `gorm:"type:varchar" json:"created_by"`

	// Updated At UTC0
	// in: int64
	UpdatedAtUtc0 int64 `gorm:"type:int8" json:"updated_at_utc0"`

	// Updated By
	// in: string
	UpdatedBy string `gorm:"type:varchar" json:"updated_by"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	time := time.Now().UnixMilli()
	tx.Statement.SetColumn("Id", uuid)
	tx.Statement.SetColumn("CreatedAtUtc0", time)
	tx.Statement.SetColumn("UpdatedAtUtc0", time)
	return nil
}

func (base *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	time := time.Now().UnixMilli()
	tx.Statement.SetColumn("UpdatedAtUtc0", time)
	return nil
}
