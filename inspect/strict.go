package inspect

import (
	"github.com/4rcode/gadget"
	"github.com/4rcode/gnomic"
)

// StrictContext is responsible for reporting an error and for failing the
// current test immediately.
type StrictContext interface {
	_context

	// Fatalf prints a formatted error message and fails the current test
	// immediately.
	Fatalf(string, ...interface{})
}

// WithStrict returns a new Inspector based on a StrictContext.
func WithStrict(context StrictContext, options ...gnomic.Option) gadget.Inspector {
	return with(context, func() _print {
		return context.Fatalf
	}, options...)
}
