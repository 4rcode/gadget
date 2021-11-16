package inspect

import (
	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/internal"
	"github.com/4rcode/gadget/settings"
)

type reporter struct {
	_context
	print _print
}

func (r *reporter) Report(value interface{}, err error, options gadget.Options) {
	value = internal.Print(options, value)
	template := settings.Template(`

provided: %[1]v
expected: %[2]v

`)

	options.ApplyToAll(&template)

	r.Helper()
	r.print(string(template), value, err, "")
}
