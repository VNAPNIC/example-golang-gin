package common

import (
	"github.com/go-playground/validator/v10"
)

var (
	Validate *validator.Validate
)

func InitValidate() {
	Validate = validator.New()
}

func CheckBindStructParameter(s interface{}) error {
	err := Validate.Struct(s)
	if err != nil {
		return err
	}
	return nil
}
