package really

import (
	"fmt"
	"reflect"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/internal"
	"github.com/4rcode/gadget/settings"
)

// Like TODO
func Like(target interface{}) gadget.DetectorFunc {
	return func(value interface{}, options gadget.Options) error {
		valueType := reflect.TypeOf(value)
		targetType := reflect.TypeOf(target)

		if !internal.IsFake(value) &&
			(valueType == targetType || (valueType != nil &&
				targetType != nil &&
				valueType.ConvertibleTo(targetType))) {
			return nil
		}

		format := settings.LikeFormat("%[1]v instance")

		options.ApplyToAll(&format)

		return fmt.Errorf(string(format), targetType, "")
	}
}
