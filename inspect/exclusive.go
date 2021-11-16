package inspect

import (
	"github.com/4rcode/gadget"
	"github.com/4rcode/gnomic"
)

// ExclusiveContext is responsible for reporting an error and for skipping the
// current test.
type ExclusiveContext interface {
	_context

	// Skipf prints a formatted error message and skips the current test.
	Skipf(string, ...interface{})
}

// WithExclusive returns a new Inspector based on a ExclusiveContext.
func WithExclusive(context ExclusiveContext, options ...gnomic.Option) gadget.Inspector {
	return with(context, func() _print {
		return context.Skipf
	}, options...)
}
