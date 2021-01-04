package api

import (
	"github.com/gin-gonic/gin"
	"go-cms/internal/app"
	"go-cms/internal/pkg"
)


func Login(c *gin.Context)  {
	var p app.LoginStru

	err :=pkg.Verify(&p,c)
	if err!=nil {
		return
	}
	p.Login(c)
	//param := schema.User{Username:c.Param("username")}


}