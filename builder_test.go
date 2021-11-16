package gadget_test

import (
	"testing"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/inspect"
	"github.com/4rcode/gadget/mock"
)

func TestBuilderFunc(t *testing.T) {
	is := inspect.With(t)

	mocked := mock.Builder("=> ")
	builder := gadget.BuilderFunc(mocked.Build)
	detector := builder.Build(1, "2", 3.4)

	is(detector, "=> 123.4")
}
