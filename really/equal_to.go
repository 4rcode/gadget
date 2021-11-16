package really

import (
	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/internal"
)

// EqualTo TODO
func EqualTo(value interface{}, values ...interface{}) gadget.Detector {
	return internal.MultiDetector{}.Values(value, values...)
}
