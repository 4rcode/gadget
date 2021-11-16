package inspect

import (
	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/internal"
	"github.com/4rcode/gadget/settings"
	"github.com/4rcode/gnomic"
)

type _context interface {
	Helper()
}

// New returns a new Inspector, configured with the provided options.
func New(options ...gnomic.Option) gadget.Inspector {
	return with(nil, nil, options...)
}

func with(context _context, print func() _print, options ...gnomic.Option) gadget.Inspector {
	var inspector internal.Inspector

	if context != nil {
		inspector.Options = inspector.Append(settings.Reporters{
			&reporter{context, print()},
		})
	}

	inspector.Options = inspector.Append(options...)

	return inspector.Inspect
}
