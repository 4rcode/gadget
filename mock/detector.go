package mock

import (
	"errors"
	"fmt"

	"github.com/4rcode/gadget"
)

type Detector string

func (d Detector) Detect(value interface{}, options gadget.Options) error {
	if string(d) == fmt.Sprint(value) {
		return nil
	}

	return errors.New(string(d))
}
