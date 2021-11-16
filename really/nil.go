package really

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/4rcode/gadget"
)

const (
	min = reflect.Array
	max = reflect.String
)

// Nil TODO
func Nil(arguments ...interface{}) gadget.DetectorFunc {
	return func(value interface{}, options gadget.Options) error {
		if value == nil {
			return nil
		}

		reference := reflect.ValueOf(value)
		kind := reference.Kind()

		if min < kind && kind < max && reference.IsNil() {
			return nil
		}

		if len(arguments) < 1 {
			arguments = append(arguments, nil)
		}

		return errors.New(fmt.Sprint(arguments...))
	}
}
