package really

import (
	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/internal"
)

// Not is responsible for detecting when a value is NOT detected by any
// Detector. Both value and values can be either a Detector or any other target
// value.
func Not(value interface{}, values ...interface{}) gadget.Detector {
	return internal.MultiDetector{
		Breaker: true, Inverted: true,
	}.Values(value, values...)
}
