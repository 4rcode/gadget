package really

import (
	"errors"

	"github.com/4rcode/gadget"
)

// InPanic TODO
func InPanic(receivers ...*interface{}) gadget.DetectorFunc {
	return func(value interface{}, options gadget.Options) error {
		f, ok := value.(func())

		if !ok || f == nil {
			return errors.New("func()")
		}

		defer func() {
			value := recover()

			for _, receiver := range receivers {
				if receiver != nil {
					*receiver = value
				}
			}
		}()

		f()

		return errors.New("panic")
	}
}
