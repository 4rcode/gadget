package gadget_test

import (
	"testing"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/inspect"
	"github.com/4rcode/gadget/mock"
)

func TestReporterFunc(t *testing.T) {
	is := inspect.With(t)

	mocked := mock.Reporter("=> ")
	detector := mock.Detector("b")
	reporter := gadget.ReporterFunc(mocked.Report)

	reporter.Report("a", detector.Detect("a", nil), nil)

	is(mocked, "=> a b")
}
