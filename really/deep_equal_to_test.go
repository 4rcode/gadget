package really_test

import (
	"testing"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/inspect"
	. "github.com/4rcode/gadget/really"
	"github.com/4rcode/gnomic"
)

func TestDeepEqualTo(t *testing.T) {
	is := inspect.With(t)

	var opts gnomic.Options

	is(DeepEqualTo(gadget.Fake{}).Detect(gadget.Fake{}, opts), String("gadget.Fake{}"))
	is(DeepEqualTo([]int{1, 2, 3}).Detect([]int{1, 2, 3}, opts), Nil())
	is(DeepEqualTo(struct{ int }{}).Detect(struct{ int }{}, opts), Nil())
	is(DeepEqualTo(nil).Detect([]struct{}(nil), opts), String("<nil>"))
}
