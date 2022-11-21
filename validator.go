package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	zh_translation "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"strings"
)

type Register struct {
	Name string `json:"name"  binding:"required,min=3,max=10"`
	Age  int    `json:"age" binding:"required,gt=18"`
}

var trans ut.Translator

func initTranslator(language string) error {
	//转换成go-playground的validator
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		//创建翻译器
		zhT := zh.New()
		enT := en.New()

		//创建通用翻译器
		//第一个参数是备用语言，后面的是应当支持的语言
		uni := ut.New(enT, enT, zhT)

		//从通过中获取指定语言翻译器
		trans, ok = uni.GetTranslator(language)
		if !ok {
			return fmt.Errorf("not found translator %s", language)
		}

		//绑定到gin的验证器上，对binding的tag进行翻译
		switch language {
		case "zh":
			err := zh_translation.RegisterDefaultTranslations(validate, trans)
			if err != nil {
				return err
			}
		default:
			err := en_translation.RegisterDefaultTranslations(validate, trans)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func remove(errors map[string]string) map[string]string {
	result := map[string]string{}
	for key, value := range errors {
		result[key[strings.Index(key, ".")+1:]] = value
	}
	return result
}

func main() {

	err := initTranslator("zh")
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		r := &Register{}
		err := c.ShouldBind(r)
		if err != nil {
			if errors, ok := err.(validator.ValidationErrors); ok {
				translate := errors.Translate(trans)
				//调用指定翻译器进行翻译
				c.JSON(http.StatusBadRequest, gin.H{
					"error1": translate,
					"error2": remove(translate),
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.Run(":9999")
}
