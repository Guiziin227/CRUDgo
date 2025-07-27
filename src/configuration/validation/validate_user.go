package validation

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/guiziin227/CRUDgo/src/configuration/c_err"
)

var (
	Validade = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *c_err.CErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return c_err.NewBadRequestErr("Invalid field type: ")
	} else if errors.As(validation_err, &jsonValidationErr) {
		errorsCauses := []c_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := c_err.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			}
			errorsCauses = append(errorsCauses, cause)
		}

		return c_err.NewBadRequestValidationErr("Some fields are invalid", errorsCauses)
	} else {
		return c_err.NewBadRequestErr("Error trying to convert fields")
	}
}
