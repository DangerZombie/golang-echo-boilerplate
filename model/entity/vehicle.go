package entity

import "go-echo/model/base"

type Vehicle struct {
	base.BaseModel

	// Brand
	// in: string
	Brand string `gorm:"type:varchar" json:"brand"`

	// Name
	// in: string
	Name string `gorm:"type:varchar" json:"name"`

	// Vehicle Number
	// in: string
	VehicleNumber string `gorm:"type:varchar" json:"vehicle_number"`

	// Color
	// in: string
	Color string `gorm:"type:varchar" json:"color"`

	// Is Available
	// in: boolean
	IsAvailable bool `gorm:"type:boolean" json:"is_available"`

	// Expired
	// in: int64
	Expired int64 `gorm:"type:bigint" json:"expired"`
}
