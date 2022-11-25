package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

type MyStruct struct {
	String string `validate:"is-awesome"`
}

// ValidateMyVal implements validator.Func
func ValidateMyVal(fl validator.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func TestValidator(t *testing.T) {
	validate = validator.New()
	err := validate.RegisterValidation("is-awesome", ValidateMyVal)
	s := MyStruct{String: "awesome"}
	err = validate.Struct(s)
	if err != nil {
		fmt.Printf("Err:%s", err.Error())
	}
	s.String = "awesome1"
	err = validate.Struct(s)
	if err != nil {
		fmt.Printf("Err:%s", err.Error())
	}
}
