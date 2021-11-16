package really_test

import (
	"testing"

	"github.com/4rcode/gadget/inspect"
	"github.com/4rcode/gadget/mock"
	. "github.com/4rcode/gadget/really"
	"github.com/4rcode/gnomic"
)

func TestEither(t *testing.T) {
	is := inspect.With(t)

	var opts gnomic.Options

	is(Either(
		mock.Detector("that's"), mock.Detector("wrong"),
	).Detect("abc", opts), String("that's | wrong"))

	is(Either(
		mock.Detector("abc"), mock.Detector("wrong"),
	).Detect("abc", opts), Nil())

	is(Either(
		mock.Detector("abc"), mock.Detector("abc"),
	).Detect("abc", opts), Nil())
}
