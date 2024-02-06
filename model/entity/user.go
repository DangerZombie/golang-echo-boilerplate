package entity

import (
	"go-echo/model/base"
)

type User struct {
	base.BaseModel

	Username string  `gorm:"type:varchar" json:"-"`
	Password string  `gorm:"type:varchar" json:"-"`
	Status   string  `gorm:"type:varchar" json:"-"`
	Nickname string  `gorm:"type:varchar" json:"-"`
	Roles    []*Role `gorm:"many2many:user_role"`
}
