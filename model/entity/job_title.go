package entity

import "go-echo/model/base"

type JobTitle struct {
	base.BaseModel

	Name        string `gorm:"type:varchar" json:"name"`
	Description string `gorm:"type:varchar" json:"description"`
}
