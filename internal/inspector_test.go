package internal_test

import (
	"testing"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/inspect"
	"github.com/4rcode/gadget/internal"
	"github.com/4rcode/gadget/mock"
	"github.com/4rcode/gadget/settings"
	"github.com/4rcode/gnomic"
)

func TestInspector(t *testing.T) {
	is := inspect.With(t)

	type opts []gnomic.Option

	type testData []struct {
		value     interface{}
		arguments []interface{}
		bool
		string
	}

	values := func(values ...interface{}) []interface{} {
		return values
	}

	var reporter mock.Reporter

	for testCaseId, testCase := range []struct {
		options opts
		testData
	}{
		{nil, testData{
			{nil, nil, false, ""},
			{nil, values(nil), true, ""},
			{nil, values(123), false, ""},
			{123, nil, false, ""},
			{123, values(nil), false, ""},
			{123, values(123), true, ""},
			{true, nil, true, ""},
			{true, values(nil), false, ""},
			{true, values(123), false, ""},
			{"abc", nil, false, ""},
			{"abc", values(nil), false, ""},
			{"abc", values(123), false, ""},
		}},
		{opts{
			settings.Format(">%#[1]v<"),
			settings.Reporters{&reporter, nil},
			settings.Builder{gadget.BuilderFunc(
				func(i ...interface{}) gadget.Detector {
					return nil
				}),
			}}, testData{
			{nil, nil, false, "#<nil> >true<"},
			{nil, values(nil), false, "#<nil> >true<"},
			{nil, values(123), false, "#<nil> >true<"},
			{123, nil, false, "#123 >true<"},
			{123, values(nil), false, "#123 >true<"},
			{123, values(123), false, "#123 >true<"},
			{true, nil, true, ""},
			{true, values(nil), true, ""},
			{true, values(123), true, ""},
			{"abc", nil, false, "#abc >true<"},
			{"abc", values(nil), false, "#abc >true<"},
			{"abc", values(123, settings.Options{settings.Format("=> %[1]v")}), false, "#abc => true"},
		}},
		{opts{
			settings.Fallback{mock.Detector("abc")},
			settings.Reporters{&reporter, nil},
		}, testData{
			{nil, nil, false, "#<nil> abc"},
			{nil, values(nil), true, ""},
			{nil, values(123), false, "#<nil> 123"},
			{123, nil, false, "#123 abc"},
			{123, values(nil), false, "#123 <nil>"},
			{123, values(123), true, ""},
			{true, nil, false, "#true abc"},
			{true, values(nil), false, "#true <nil>"},
			{true, values(123), false, "#true 123"},
			{"abc", nil, true, ""},
			{"abc", values(nil), false, "#abc <nil>"},
			{"abc", values(123), false, "#abc 123"},
		}},
	} {
		inspector := internal.Inspector{
			gnomic.Options(testCase.options),
		}

		for testDataId, testData := range testCase.testData {
			t.Log(testCaseId+1, testDataId+1)

			reporter = ""

			is(inspector.Inspect(testData.value, testData.arguments...), testData.bool)
			is(reporter, testData.string)
		}
	}
}
