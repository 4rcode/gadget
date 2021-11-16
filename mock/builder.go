package mock

import (
	"fmt"

	"github.com/4rcode/gadget"
)

type Builder string

func (b Builder) Build(a ...interface{}) gadget.Detector {
	return Detector(string(b) + fmt.Sprint(a...))
}
