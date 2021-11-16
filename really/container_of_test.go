package really_test

import (
	"testing"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/inspect"
	. "github.com/4rcode/gadget/really"
	"github.com/4rcode/gnomic"
)

func TestContainerOf(t *testing.T) {
	is := inspect.With(t)

	var opts gnomic.Options

	is(ContainerOf(gadget.Fake{}).Detect(gadget.Fake{}, opts), String("not implemented"))
	is(ContainerOf([]int{1, 2, 3}).Detect([]int{1, 2, 3}, opts), String("not implemented"))
	is(ContainerOf(struct{ int }{}).Detect(struct{ int }{}, opts), String("not implemented"))
	is(ContainerOf(nil).Detect([]struct{}(nil), opts), String("not implemented"))
}
