package schema

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null;"`
	Sort int8 `gorm:"default:0"`
}
