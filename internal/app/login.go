package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


type LoginStru struct {
	Username string `form:"username" binding:"required,logincheck"`
	Password string `form:"password" binding:"required"`
}

func LoginFormatCheck(f1 validator.FieldLevel) bool {
	return false
}
func (p *LoginStru) Login(c *gin.Context) {
	p.Password="ss"+p.Password
	//pkg.Dao.Create(&p)
	//c.JSON(http.StatusOK,pkg.ResJsonData{Code:200,Msg:"test"})
}