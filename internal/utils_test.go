package internal_test

import (
	"testing"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/inspect"
	"github.com/4rcode/gadget/internal"
	"github.com/4rcode/gadget/settings"
	"github.com/4rcode/gnomic"
)

func TestComparable(t *testing.T) {
	is := inspect.With(t)

	is(internal.Comparable(nil))
	is(internal.Comparable(123))
	is(internal.Comparable(" "))
	is(internal.Comparable((*struct{})(nil)))
	is(internal.Comparable([]struct{}{}), false)
	is(internal.Comparable(map[struct{}]struct{}{}), false)
}

func TestConvert(t *testing.T) {
	is := inspect.With(t)

	is(internal.Convert(1.2, 3), 1)
	is(internal.Convert(123, nil), 123)
	is(internal.Convert(nil, 123), nil)
	is(internal.Convert(struct{ int }{3}, struct{ int }{7}), struct{ int }{3})
	is(internal.Convert("1.2", 3), "1.2")
	is(internal.Convert(struct{ int }{}, struct{ string }{}), struct{ int }{})
}

func TestIsFake(t *testing.T) {
	is := inspect.With(t)

	type fake gadget.Fake

	is(internal.IsFake(gadget.Fake{}))
	is(internal.IsFake(nil), false)
	is(internal.IsFake(123), false)
	is(internal.IsFake("abc"), false)
	is(internal.IsFake(fake{}), false)
	is(internal.IsFake(struct{}{}), false)
	is(internal.IsFake(&gadget.Fake{}), false)
}

func TestPrint(t *testing.T) {
	is := inspect.With(t)

	var cfg gnomic.Options

	cfg1 := gnomic.Options{settings.Format("%[1]T")}
	cfg2 := gnomic.Options{settings.Format("abc%[2]v")}

	is(internal.Print(cfg, func() {}), "github.com/4rcode/gadget/internal_test.TestPrint.func1")
	is(internal.Print(cfg1, func() {}), "github.com/4rcode/gadget/internal_test.TestPrint.func2")
	is(internal.Print(cfg2, func() {}), "github.com/4rcode/gadget/internal_test.TestPrint.func3")
	is(internal.Print(cfg, 123), "123")
	is(internal.Print(cfg1, 123), "int")
	is(internal.Print(cfg2, 123), "abc")
}
