package entity

type User struct {
	Username string `gorm:"type:varchar" json:"-"`
	Password string `gorm:"type:varchar" json:"-"`
}
