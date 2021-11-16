package really_test

import (
	"testing"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/inspect"
	. "github.com/4rcode/gadget/really"
	"github.com/4rcode/gnomic"
)

func TestZero(t *testing.T) {
	is := inspect.With(t)

	var opts gnomic.Options

	is(Zero().Detect(gadget.Fake{}, opts), String("gadget.Fake{}"))
	is(Zero().Detect(struct{ int }{1}, opts), String("struct { int }{int:0}"))
	is(Zero().Detect("abc", opts), String("\"\""))
	is(Zero().Detect(123, opts), String("0"))
	is(Zero().Detect(nil, opts), Nil())
	is(Zero().Detect("", opts), Nil())
}
