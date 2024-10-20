package entity

import "go-echo/model/base"

type Teacher struct {
	base.BaseModel

	Status     string `gorm:"type:varchar" json:"status"`
	Experience int    `gorm:"type:int" json:"experience"`
	Degree     string `gorm:"type:varchar" json:"degree"`

	JobTitleId string   `json:"job_title_id"`
	JobTitle   JobTitle `gorm:"foreignKey:JobTitleId" json:"job_title"`

	UserId string `json:"user_id"`
	User   User   `gorm:"foreignKey:UserId" json:"user"`
}
