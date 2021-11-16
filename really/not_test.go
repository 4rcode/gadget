package really_test

import (
	"testing"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/inspect"
	"github.com/4rcode/gadget/mock"
	. "github.com/4rcode/gadget/really"
	"github.com/4rcode/gnomic"
)

func TestNot(t *testing.T) {
	is := inspect.With(t)

	var opts gnomic.Options

	is(Not(nil).Detect("", opts), Nil())
	is(Not("").Detect("", opts), String("not \"\""))
	is(Not(True("really?")).Detect(false, opts), Nil())
}

func TestNoneWithValues(t *testing.T) {
	is := inspect.With(t)

	var opts gnomic.Options

	for ti, td := range []struct {
		gadget.Detector
		string
	}{
		{Not(""), "<nil>"},
		{Not("value"), "not \"value\""},
		{Not("different", "value"), "not \"value\""},
		{Not(mock.Detector("not"), mock.Detector("again")), "<nil>"},
	} {
		t.Log(ti+1, td)

		is(td.Detector.Detect("value", opts), String(td.string))
	}
}
