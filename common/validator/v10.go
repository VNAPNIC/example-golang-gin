package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"sync"
)

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &DefaultValidator{}

// ValidateStruct Validation is performed if the received type is a struct or a pointer to a struct.
func (v *DefaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		// Returns InvalidValidationError if an illegal value is passed, otherwise returns nil.
		///If err != nil is returned, the errors array can be accessed via err.(validator.ValidationErrors).
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

// Engine Returns the underlying validation engine that supports the `StructValidator` implementation
func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("validate")
		// v8 version, v8 version uses "binding"
		// v.validate.SetTagName("binding")
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
