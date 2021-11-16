package settings_test

import (
	"testing"

	"github.com/4rcode/gadget/inspect"
	"github.com/4rcode/gadget/mock"
	"github.com/4rcode/gadget/settings"
	"github.com/4rcode/gnomic"
)

func TestSettings(t *testing.T) {
	is := inspect.With(t)

	builder1 := mock.Builder("123")
	builder2 := mock.Builder("abc")

	reporter1 := mock.Reporter("123")
	reporter2 := mock.Reporter("abc")

	opts := struct {
		settings.Builder
		settings.Fallback
		settings.Format
		settings.LikeFormat
		settings.NegativeFormat
		settings.StringFormat
		settings.Separator
		settings.Template
		settings.Reporters
	}{
		settings.Builder{builder1},
		settings.Fallback{"123", nil},
		"123", "123", "123", "123", "123", "123",
		settings.Reporters{&reporter1, nil},
	}

	gnomic.Options{
		settings.Builder{builder2},
		settings.Fallback{"x", "y", "z"},
		settings.Fallback{},
		settings.Fallback{"a"},
		settings.Fallback{"b", "c"},
		settings.Format("abc"),
		settings.LikeFormat("abc"),
		settings.NegativeFormat("abc"),
		settings.Reporters{&reporter1, nil},
		settings.Reporters{},
		settings.Reporters{&reporter2},
		settings.StringFormat("abc"),
		settings.Separator("abc"),
		settings.Template("abc"),
	}.ApplyTo(&opts)

	is(opts.Builder[0], builder2)

	if is(len(opts.Fallback), 3) {
		is(opts.Fallback[0], "a")
		is(opts.Fallback[1], "b")
		is(opts.Fallback[2], "c")
	}

	is(opts.Format, "abc")
	is(opts.LikeFormat, "abc")
	is(opts.NegativeFormat, "abc")

	if is(len(opts.Reporters), 1) {
		is(opts.Reporters[0], &reporter2)
	}

	is(opts.StringFormat, "abc")
	is(opts.Separator, "abc")
	is(opts.Template, "abc")
}
