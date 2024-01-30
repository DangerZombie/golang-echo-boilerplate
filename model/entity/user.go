package entity

type User struct {
	Id            string `gorm:"type:uuid" json:"-"`
	Username      string `gorm:"type:varchar" json:"-"`
	Password      string `gorm:"type:varchar" json:"-"`
	Status        string `gorm:"type:varchar" json:"-"`
	Nickname      string `gorm:"type:varchar" json:"-"`
	CreatedAtUtc0 int64  `gorm:"type:int8" json:"-"`
	CreatedBy     string `gorm:"type:varchar" json:"-"`
	UpdatedAtUtc0 int64  `gorm:"type:int8" json:"-"`
	UpdatedBy     string `gorm:"type:varchar" json:"-"`
}
