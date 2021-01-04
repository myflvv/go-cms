package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


type LoginStru struct {
	Username string `form:"username" binding:"required" label:"用户名"`
	Password string `form:"password" binding:"required" label:"密码"`
}

func (p *LoginStru) Login(c *gin.Context) {
	p.Password="ss"+p.Password
	fmt.Print("sss")
	//pkg.Dao.Create(&p)
	//c.JSON(http.StatusOK,pkg.ResJsonData{Code:200,Msg:"test"})
}