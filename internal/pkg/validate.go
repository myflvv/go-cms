package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go-cms/internal/app"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

func InitVali()  {
	uni := ut.New(zh.New())
	Trans,_=uni.GetTranslator("zh")
	if v,ok:=binding.Validator.Engine().(*validator.Validate); ok {
		_=zh_translations.RegisterDefaultTranslations(v,Trans)
		_=v.RegisterValidation("logincheck",app.LoginFormatCheck)

		//_=v.RegisterTranslation("logincheck", Trans, func(ut ut.Translator) error {
		//	return ut.Add("logincheck", "{0}不能早于当前时间或{1}格式错误!", true)
		//}, func(ut ut.Translator, fe validator.FieldError) string {
		//	t, _ := ut.T("logincheck", fe.Field(), fe.Field())
		//	return t
		//})
	}
}

func Request(c *gin.Context,request interface{}) error {
	err:=c.ShouldBind(request)
	if err!=nil {
		RespJson(c,&ResponseData{
			Code:CODE_FAULT,
			Msg:err.Error(),
		})
		return err
	}
	return nil
}