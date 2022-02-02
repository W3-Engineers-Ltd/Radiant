package validation

import (
	"reflect"

	"github.com/W3-Engineers-Ltd/Radiant/core/validation"
)

const (
	// ValidTag struct tag
	ValidTag = validation.ValidTag

	LabelTag = validation.LabelTag
)

var ErrInt64On32 = validation.ErrInt64On32

// CustomFunc is for custom validate function
type CustomFunc func(v *Validation, obj interface{}, key string)

// AddCustomFunc Add a custom function to validation
// The name can not be:
//   Clear
//   HasErrors
//   ErrorMap
//   Error
//   Check
//   Valid
//   NoMatch
// If the name is same with exists function, it will replace the origin valid function
func AddCustomFunc(name string, f CustomFunc) error {
	return validation.AddCustomFunc(name, func(v *validation.Validation, obj interface{}, key string) {
		f((*Validation)(v), obj, key)
	})
}

// ValidFunc Valid function type
type ValidFunc validation.ValidFunc

// Funcs Validate function map
type Funcs validation.Funcs

// Call validate values with named type string
func (f Funcs) Call(name string, params ...interface{}) (result []reflect.Value, err error) {
	return (validation.Funcs(f)).Call(name, params...)
}
