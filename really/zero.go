package really

import (
	"errors"
	"reflect"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/internal"
)

// Zero TODO
func Zero() gadget.DetectorFunc {
	return func(value interface{}, options gadget.Options) error {
		if value == nil {
			return nil
		}

		reference := reflect.ValueOf(value)

		if !internal.IsFake(value) && reference.IsZero() {
			return nil
		}

		nature := reference.Type()
		zero := reflect.Zero(nature)

		return errors.New(internal.Print(options, zero))
	}
}
