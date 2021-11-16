package inspect_test

import (
	"testing"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/inspect"
	"github.com/4rcode/gadget/mock"
	"github.com/4rcode/gadget/settings"
)

func TestContexts(t *testing.T) {
	is := inspect.With(t)

	var context mock.Context

	expected := `##

provided: 123
expected: true

`

	for testCaseId, inspector := range []gadget.Inspector{
		inspect.New(nil, settings.Format("%[1]v")),
		inspect.With(&context),
		inspect.WithExclusive(&context),
		inspect.WithLenient(&context),
		inspect.WithStrict(&context),
	} {
		t.Log(testCaseId + 1)

		context = ""

		is(inspector(123), false)

		if testCaseId > 0 {
			is(context, expected)
		}
	}
}
