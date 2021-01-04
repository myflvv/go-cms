package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CODE_SUCCESS int = 0
	CODE_FAULT int = 1
)

type ResponseData struct {
	Code int `json:"code"`
	Msg interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func RespJson(res *ResponseData,c *gin.Context)  {
	c.JSON(http.StatusOK,res)
}