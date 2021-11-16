package really_test

import (
	"testing"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/inspect"
	. "github.com/4rcode/gadget/really"
	"github.com/4rcode/gnomic"
)

func TestLike(t *testing.T) {
	is := inspect.With(t)

	var opts gnomic.Options

	is(Like(gadget.Fake{}).Detect(gadget.Fake{}, opts), String("gadget.Fake instance"))
	is(Like(nil).Detect(123, opts), String("<nil> instance"))
	is(Like("").Detect("abc", opts), Nil())
	is(Like(0.0).Detect(0, opts), Nil())
	is(Like(0).Detect(123, opts), Nil())
}
