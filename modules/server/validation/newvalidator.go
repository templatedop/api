package validation

import (
	"errors"
	"fmt"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	ut "github.com/templatedop/universal-translator-master"
)


type CustomValidation struct {
	Tag     string
	Func    validator.Func
	Message string
	Code    string
}

type IValidatorService interface {
	RegisterCustomValidation(tag string, fn validator.Func, message string, code string) error
	ValidateStruct(s interface{}) error	
}

type ValidatorService struct {
	validate          *validator.Validate
	trans             ut.Translator
	customValidations map[string]CustomValidation	
}


func NewValidatorService() (IValidatorService, error) {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	validate.RegisterTagNameFunc(getStructFieldName)
	err := en_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return nil, err
	}

	return &ValidatorService{
		validate:          validate,
		trans:             trans,
		customValidations: make(map[string]CustomValidation),
		
	}, nil
}

func (vs *ValidatorService) RegisterCustomValidation(tag string, fn validator.Func, message string, code string) error {
	if tag == "" {
		return errors.New("validation tag cannot be empty")
	}
	if fn == nil {
		return errors.New("validation function cannot be nil")
	}

	if _, exists := vs.customValidations[tag]; exists {
		return fmt.Errorf("validation tag '%s' is already registered", tag)
	}

	err := vs.validate.RegisterValidation(tag, fn)
	if err != nil {
		return fmt.Errorf("failed to register validation for tag '%s': %v", tag, err)
	}
	vs.customValidations[tag] = CustomValidation{Tag: tag, Func: fn, Message: message, Code: code}
	return nil
}

func (vs *ValidatorService) ValidateStruct(s interface{}) error {

	//var fieldErrors validation.FieldErrors
	err := vs.validate.Struct(s)
	if err != nil {
		var fieldErrors []FieldError
		var validatorErrors validator.ValidationErrors

		errors.As(err, &validatorErrors)


		// for _, err := range validatorErrors {
		// 	fieldErrors = append(fieldErrors, FieldError{
		// 		FailedField: err.Field(),
		// 		Tag:         err.Tag(),
		// 		Value:       err.Value(),
		// 		Message:     getTagMessage(err),
		// 	})
		// }

		for _, e := range validatorErrors {

			tag := e.Tag()
			if cv, ok := vs.customValidations[tag]; ok {
				fieldErrors = append(fieldErrors, FieldError{
					FailedField: e.Field(),
					//Tag:         e.Tag(),
					//Value:       e.Value(),
					Message:     cv.Message,
				})

			} else {
				fieldErrors = append(fieldErrors, FieldError{
					FailedField: e.Field(),
					//Tag:         e.Tag(),
					//Value:       e.Value(),
					Message:     e.Translate(vs.trans),
				})

			}

		}

		return &Error{
			msg:  "validation error",
			errs: fieldErrors,
		}

		//err := response.Error("validation error", fieldErrors...)

		//return err

	}
	return nil
}
