package schema

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Ip string `gorm:"type:varchar(50)"`
	Website string `gorm:"type:int(1)"`
}
