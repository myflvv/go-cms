package app

import (
	"github.com/gin-gonic/gin"
	"go-cms/schema"
)

type User struct {
	*schema.User
}


func (p *User) Loginss(c *gin.Context) {
	p.Password="ss"+p.Password
	//pkg.Dao.Create(&p)
	//c.JSON(http.StatusOK,pkg.ResJsonData{Code:200,Msg:"test"})
}