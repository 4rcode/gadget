package really

import (
	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/internal"
)

// Either TODO
func Either(value interface{}, values ...interface{}) gadget.Detector {
	return internal.MultiDetector{
		Breaker: true,
	}.Values(value, values...)
}
