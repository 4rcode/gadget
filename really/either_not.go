package really

import (
	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/internal"
)

// EitherNot TODO
func EitherNot(value interface{}, values ...interface{}) gadget.Detector {
	return internal.MultiDetector{
		Breaker: false, Inverted: true,
	}.Values(value, values...)
}
