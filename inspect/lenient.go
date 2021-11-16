package inspect

import (
	"github.com/4rcode/gadget"
	"github.com/4rcode/gnomic"
)

// LenientContext is responsible for reporting an error without side effects on
// the current test.
type LenientContext interface {
	_context

	// Logf prints a formatted error message.
	Logf(format string, arguments ...interface{})
}

// WithLenient returns a new Inspector based on a LenientContext.
func WithLenient(context LenientContext, options ...gnomic.Option) gadget.Inspector {
	return with(context, func() _print {
		return context.Logf
	}, options...)
}
