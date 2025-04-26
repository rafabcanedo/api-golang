package validation

import (
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/rafabcanedo/api-golang/src/configuration/rest_errors"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl ut.Translator
)

// Validation Translate Rules
func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(
	validation_err error,
) *rest_errors.RestErrors {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	// & in jsonError and jsonValidationError, because it's warning indicate this params
	if errors.As(validation_err, &jsonErr) {
		return rest_errors.NewBadRequestError("invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_errors.Causes{}

		// For each error, it create a Causes structure with Message and which field has error
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_errors.Causes{
				Message: e.Translate(transl),
				Field: e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_errors.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_errors.NewBadRequestError("Error trying to convert fields")
	}
}