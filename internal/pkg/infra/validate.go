package infra

import (
	"log"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	tzh "github.com/go-playground/validator/v10/translations/zh"
	z "go.uber.org/zap"
)

var (
	validate   *validator.Validate
	translator ut.Translator
)

func init() {
	// Create customize validator
	validate = validator.New()
	cn := zh.New()
	uni := ut.New(cn, cn)
	var found bool

	// Get Translator
	translator, found = uni.GetTranslator("zh")
	if found {
		if err := tzh.RegisterDefaultTranslations(validate, translator); err != nil {
			log.Fatal("Set default translation fail", z.Reflect("msg", err.Error()))
		}
	} else {
		log.Fatal("Not found translator")
	}
}
