package internal

import (
	"errors"
	"fmt"

	"github.com/4rcode/gadget"
)

const comparableFormat = "a comparable target value: try DeepEqualTo(%#v)"

func equalTo(target interface{}) gadget.DetectorFunc {
	return func(value interface{}, options gadget.Options) error {
		if !Comparable(target) {
			return fmt.Errorf(comparableFormat, target)
		}

		if !IsFake(value) && Convert(value, target) == target {
			return nil
		}

		return errors.New(Print(options, target))
	}
}
