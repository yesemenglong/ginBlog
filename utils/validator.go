package utils

import (
	"fmt"

	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

func Validator(data interface{}) (err error) {
	vali := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err = zhTrans.RegisterDefaultTranslations(vali, trans)
	if err != nil {
		fmt.Println("err:", err)
	}
	err = vali.Struct(data)
	return err
}
