package really

import (
	"errors"

	"github.com/4rcode/gadget"
)

// ContainerOf TODO
func ContainerOf(value interface{}, values ...interface{}) gadget.DetectorFunc {
	return func(value interface{}, options gadget.Options) error {
		return errors.New("not implemented")
	}
}
