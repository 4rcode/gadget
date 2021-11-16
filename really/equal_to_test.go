package really_test

import (
	"testing"

	"github.com/4rcode/gadget/inspect"
	"github.com/4rcode/gadget/mock"
	. "github.com/4rcode/gadget/really"
	"github.com/4rcode/gnomic"
)

func TestEqualTo(t *testing.T) {
	is := inspect.With(t)

	var opts gnomic.Options

	is(EqualTo(
		mock.Detector("that's"), mock.Detector("wrong"),
	).Detect("", opts), String("that's"))

	is(EqualTo(
		mock.Detector(""), mock.Detector("wrong"),
	).Detect("", opts), String("wrong"))

	is(EqualTo(
		mock.Detector(""), mock.Detector(""),
	).Detect("", opts), Nil())
}
