package api

import (
	"github.com/gin-gonic/gin"
	"go-cms/internal/app"
	"go-cms/internal/pkg"
)


func Login(c *gin.Context)  {
	//var res *pkg.ResponseData
	var p app.LoginStru
	//if err := c.ShouldBind(&p);err != nil {
	//	//res.Code=pkg.CODE_FAULT
	//	//res.Msg=err.Error()
	//	//res.ResData(c)
	//	errs:=err.(validator.ValidationErrors)
	//	fmt.Println(errs.Translate(pkg.Trans))
	//	pkg.RespJson(c,&pkg.ResponseData{Code:pkg.CODE_FAULT,Msg:err.Error()})
	//	return
	//}
	//var err error
	err :=pkg.Request(c,&p)
	if err!=nil {
		return
	}
	p.Login(c)
	//param := schema.User{Username:c.Param("username")}


}