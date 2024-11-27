package vldtor

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zhtranslations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	z := zh.New()
	uni = ut.New(z, z)

	trans, _ = uni.GetTranslator("zh")

	validate = validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	_ = zhtranslations.RegisterDefaultTranslations(validate, trans)

	RegisterCustomRule(trans, validate)
}

// Validator data
func Validator(data interface{}) (msg []string, isValidator bool) {
	if err := validate.Struct(data); err != nil {
		errs := err.(validator.ValidationErrors)

		for _, value := range errs.Translate(trans) {
			msg = append(msg, value)
		}
		return
	}
	return nil, true
}

// PostValidator 集合参数绑定和参数校验，传入参数需要处理成指针
func PostValidator(ctx *gin.Context, dataPtr interface{}) (msg []string, isValidator bool) {
	if err := ctx.ShouldBind(dataPtr); err != nil {
		msg = append(msg, err.Error())
		return
	}

	return Validator(dataPtr)
}
