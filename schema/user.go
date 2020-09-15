package schema

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null" form:"username" binding:"required"`
	Password string `gorm:"type:varchar(100);not null" form:"password"`
	Ip string `gorm:"type:varchar(50)"`
}
