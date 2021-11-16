package really_test

import (
	"testing"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/inspect"
	. "github.com/4rcode/gadget/really"
)

func TestTrue(t *testing.T) {
	is := inspect.With(t)

	type testData []struct {
		value interface{}
		string
	}

	for testCaseId, testCase := range []struct {
		gadget.Detector
		testData
	}{
		{True("abc"), testData{
			{nil, "abc"},
			{123, "abc"},
			{"abc", "abc"},
			{false, "abc"},
			{true, "<nil>"},
		}},
		{True("%v", 123), testData{
			{nil, "123"},
			{123, "123"},
			{"abc", "123"},
			{false, "123"},
			{true, "<nil>"},
		}}} {
		for testDataId, testData := range testCase.testData {
			t.Log(testCaseId+1, testDataId+1)

			is(testCase.Detector.Detect(testData.value, nil), String(testData.string))
		}
	}
}
