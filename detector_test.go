package gadget_test

import (
	"testing"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/inspect"
	"github.com/4rcode/gadget/mock"
	. "github.com/4rcode/gadget/really"
)

func TestDetectorFunc(t *testing.T) {
	is := inspect.With(t)

	mocked := mock.Detector("123")
	detector := gadget.DetectorFunc(mocked.Detect)

	is(detector.Detect(123, nil), Nil())
	is(detector.Detect("abc", nil), String("123"))
}

func TestFake(t *testing.T) {
	is := inspect.With(t)

	is(gadget.Fake{}.String(), "Fake")
}
