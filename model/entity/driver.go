package entity

import "go-echo/model/base"

type Driver struct {
	base.BaseModel

	// Driver Name
	// in: string
	Name string `gorm:"type:varchar" json:"name"`

	// License Number
	// in: string
	LicenseNumber string `gorm:"type:varchar" json:"license_number"`

	// Is Available
	// in: bool
	IsAvailable bool `gorm:"type:boolean" json:"is_available"`
}
