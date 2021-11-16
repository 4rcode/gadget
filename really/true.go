package really

import (
	"fmt"

	"github.com/4rcode/gadget"
)

// True TODO
func True(format string, arguments ...interface{}) gadget.DetectorFunc {
	return func(value interface{}, options gadget.Options) error {
		if value == true {
			return nil
		}

		return fmt.Errorf(format, arguments...)
	}
}
