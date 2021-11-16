package really

import (
	"errors"
	"reflect"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/internal"
)

// DeepEqualTo TODO
func DeepEqualTo(target interface{}) gadget.DetectorFunc {
	return func(value interface{}, options gadget.Options) error {
		if !internal.IsFake(value) &&
			reflect.DeepEqual(
				internal.Convert(value, target), target) {
			return nil
		}

		return errors.New(internal.Print(options, target))
	}
}
