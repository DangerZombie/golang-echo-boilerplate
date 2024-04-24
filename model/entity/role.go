package entity

type Role struct {
	Id            string  `gorm:"type:uuid;primaryKey" json:"-"`
	Name          string  `gorm:"type:varchar" json:"-"`
	Scope         string  `gorm:"type:varchar" json:"-"`
	CreatedAtUtc0 int64   `gorm:"type:int8" json:"-"`
	CreatedBy     string  `gorm:"type:varchar" json:"-"`
	UpdatedAtUtc0 int64   `gorm:"type:int8" json:"-"`
	UpdatedBy     string  `gorm:"type:varchar" json:"-"`
	Users         []*User `gorm:"many2many:user_role"`
}
