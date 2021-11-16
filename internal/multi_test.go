package internal_test

import (
	"testing"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/inspect"
	"github.com/4rcode/gadget/internal"
	. "github.com/4rcode/gadget/really"
	"github.com/4rcode/gadget/settings"
	"github.com/4rcode/gnomic"
)

func TestMultiDetector(t *testing.T) {
	is := inspect.With(t)

	type opts []gnomic.Option

	type testData []struct {
		value interface{}
		string
	}

	for testCaseId, testCase := range []struct {
		internal.MultiDetector
		options opts
		testData
	}{
		{
			internal.MultiDetector{
				Breaker: false, Inverted: false,
			}, nil, testData{
				{(*testing.T)(nil), "<nil>"},
				{nil, "*testing.T instance"},
				{123, "NIL"},
			},
		}, {
			internal.MultiDetector{
				Breaker: false, Inverted: true,
			}, opts{
				settings.LikeFormat("~ %[1]v"),
			}, testData{
				{(*testing.T)(nil), "not NIL | not ~ *testing.T"},
				{nil, "<nil>"},
				{123, "<nil>"},
			},
		}, {
			internal.MultiDetector{
				Breaker: true, Inverted: false,
			}, opts{
				settings.Format("%#[1]v"),
				settings.LikeFormat("like %[1]v"),
				settings.Separator(" or "),
			}, testData{
				{(*testing.T)(nil), "<nil>"},
				{nil, "<nil>"},
				{123, "NIL or like *testing.T"},
			},
		}, {
			internal.MultiDetector{
				Breaker: true, Inverted: true,
			}, opts{
				settings.NegativeFormat("!(%[1]v)"),
			}, testData{
				{(*testing.T)(nil), "!(NIL)"},
				{nil, "!(NIL)"},
				{123, "<nil>"},
			},
		},
	} {
		opts := gnomic.Options(testCase.options)
		detector := testCase.MultiDetector.Values(Nil("NIL"), Like(t))

		for testDataId, testData := range testCase.testData {
			t.Log(testCaseId+1, testDataId+1)

			is(detector.Detect(testData.value, opts), String(testData.string))
		}
	}
}

func TestMultiDetectorFake(t *testing.T) {
	is := inspect.With(t)

	var opts gnomic.Options

	detector := internal.MultiDetector{}.Values(gadget.Fake{})

	is(detector.Detect(gadget.Fake{}, opts), String("gadget.Fake{}"))
}

func TestMultiDetectorNotComparable(t *testing.T) {
	is := inspect.With(t)

	var opts gnomic.Options

	detector := internal.MultiDetector{}.Values([]struct{}{})
	expected := "a comparable target value: try DeepEqualTo([]struct {}{})"

	is(detector.Detect(123, opts), String(expected))
}
