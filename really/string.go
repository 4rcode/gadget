package really

import (
	"fmt"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/internal"
	"github.com/4rcode/gadget/settings"
)

// String TODO
func String(target interface{}) gadget.DetectorFunc {
	return func(value interface{}, options gadget.Options) error {
		v := fmt.Sprint(value)
		t := fmt.Sprint(target)

		if !internal.IsFake(value) && v == t {
			return nil
		}

		format := settings.StringFormat("string equality\n\n%#[1]v (provided)\n%#[2]v (expected)")

		options.ApplyToAll(&format)

		return fmt.Errorf(string(format), v, t, "")
	}
}
