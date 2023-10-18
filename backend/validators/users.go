package validators

import "github.com/go-playground/validator/v10"

var EmailValidator validator.Func = func(fl validator.FieldLevel) bool {
	//email, ok := fl.Field().Interface().(email)
	return true
}
