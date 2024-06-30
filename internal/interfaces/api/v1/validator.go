package v1

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type CustomValidator struct {
	validator  *validator.Validate
	translator ut.Translator
}

func NewValidator() *CustomValidator {
	validate := validator.New()
	translator := initTranslations(validate)
	return &CustomValidator{
		validator:  validate,
		translator: translator,
	}
}

func initTranslations(validate *validator.Validate) ut.Translator {
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	translator, _ := uni.GetTranslator("en")

	enTranslations.RegisterDefaultTranslations(validate, translator)

	// Register custom translations
	registerCustomTranslations(validate, translator)

	return translator
}

func registerCustomTranslations(validate *validator.Validate, translator ut.Translator) {
	validate.RegisterTranslation("required", translator, func(ut ut.Translator) error {
		return ut.Add("required", "{0} field is required", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
}

func (cv *CustomValidator) ValidateStruct(s interface{}) []string {
	err := cv.validator.Struct(s)
	if err == nil {
		return nil
	}
	var errorMessages []string
	for _, err := range err.(validator.ValidationErrors) {
		errorMessages = append(errorMessages, err.Translate(cv.translator))
	}
	return errorMessages
}
