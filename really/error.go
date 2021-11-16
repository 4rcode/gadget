package really

import (
	"errors"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/internal"
)

// Error TODO
func Error(target error) gadget.DetectorFunc {
	return func(value interface{}, options gadget.Options) error {
		err, ok := value.(error)

		if (ok || value == nil) && errors.Is(err, target) {
			return nil
		}

		return errors.New(internal.Print(options, target))
	}
}
