package inspect

import (
	"github.com/4rcode/gadget"
	"github.com/4rcode/gnomic"
)

// TestingContext is responsible for reporting an error and for failing the
// current test.
type TestingContext interface {
	_context

	// Errorf prints a formatted error message and fails the current test.
	Errorf(string, ...interface{})
}

// With returns a new Inspector based on a TestingContext.
func With(context TestingContext, options ...gnomic.Option) gadget.Inspector {
	return with(context, func() _print {
		return context.Errorf
	}, options...)
}
