package really_test

import (
	"testing"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/inspect"
	. "github.com/4rcode/gadget/really"
	"github.com/4rcode/gadget/settings"
	"github.com/4rcode/gnomic"
)

func TestString(t *testing.T) {
	is := inspect.With(t)

	var opts gnomic.Options

	is(String(gadget.Fake{}).Detect(gadget.Fake{}, opts), String("string equality\n\n\"Fake\" (provided)\n\"Fake\" (expected)"))
	is(String("abc").Detect(nil, opts), String("string equality\n\n\"<nil>\" (provided)\n\"abc\" (expected)"))
	is(String("abc").Detect(123, opts), String("string equality\n\n\"123\" (provided)\n\"abc\" (expected)"))
	is(String("abc").Detect("abc", opts), Nil())

	opts = opts.Append(
		settings.StringFormat("%#[1]v|%#[2]v"),
	)

	is(String("abc").Detect(123, opts), String("\"123\"|\"abc\""))
}
