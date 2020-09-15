package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-cms/internal/app"
)


func Login(c *gin.Context)  {
	var p app.User
	if err := c.ShouldBind(&p);err != nil {
		fmt.Println("username error")
		return
	}
	p.Login()
	//param := schema.User{Username:c.Param("username")}


}