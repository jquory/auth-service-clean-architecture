package exceptions

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
)

func Validate(fieldValidator interface{}) {
	validate := validator.New()
	err := validate.Struct(fieldValidator)
	if err != nil {
		var messages []map[string]interface{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]interface{}{
				"field":   err.Field(),
				"message": "This field is " + err.Tag(),
			})
		}

		jsonMassage, errJson := json.Marshal(messages)
		PanicLogging(errJson)
		panic(ValidationError{
			Message: string(jsonMassage),
		})
	}
}
