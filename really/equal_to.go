package really

import (
	"fmt"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/internal"
)

type Equal struct {
	value interface{}

	Format string
}

func (e Equal) Detect(value interface{}, options gadget.Options) error {
	return nil
}

func (e Equal) To(value interface{}) gadget.Detector {
	e.value = value
	return e
}

func _() {
	fmt.Println(Equal{}.To(3))
	equal := Equal{Format: "ffsf"}

	equal.ToValue(3)
}

// EqualTo TODO
func EqualTo(value interface{}, values ...interface{}) gadget.Detector {
	return internal.MultiDetector{}.Values(value, values...)
}
