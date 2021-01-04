package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var trans ut.Translator

func InitVali() (err error) {
	uni := ut.New(zh.New())
	trans,_=uni.GetTranslator("zh")
	if v,ok:=binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("label"),",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		//_=v.RegisterValidation("logincheck",app.LoginFormatCheck)
		err = zh_translations.RegisterDefaultTranslations(v,trans)
		return
	}
	return
}

func Verify(request interface{},c *gin.Context) error {
	_=InitVali()
	err:=c.ShouldBind(request)
	if err!=nil {
		var msg interface{}
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			msg = err.Error()
		}else {
			msg = removeTostring(errs.Translate(trans))
		}
		RespJson(&ResponseData{
			Code:CODE_FAULT,
			Msg:msg,
		},c)
		return err
	}
	return nil
}

//转换map为string,只显示第一个错误信息
func removeTostring(fields map[string]string) string {
	//只显示第一个错误
	i:=0
	var res string
	for _, err := range fields {
		if i==0 {
			res=err
		}
		i++
	}
	return res
}