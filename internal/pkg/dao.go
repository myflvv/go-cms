package pkg

import (
	"fmt"
	"go-cms/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Dao *gorm.DB

func init()  {
	var err error
	dsn:=fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.GetString("db.user"),
		config.GetString("db.pass"),
		config.GetString("db.host"),
		config.GetString("db.port"),
		config.GetString("db.name"),
		config.GetString("db.charset"),
		)
	Dao,err = gorm.Open(mysql.Open(dsn),&gorm.Config{
			NamingStrategy:schema.NamingStrategy{
				TablePrefix:config.GetString("db.table_prefix"),
				SingularTable:true,
			},
	})
	if err !=nil {
		Logger.Panic("连接数据库失败")
	}
}