package really_test

import (
	"io"

	"github.com/4rcode/gadget/dummy"
	"github.com/4rcode/gadget/inspect"
	. "github.com/4rcode/gadget/really"
)

var t dummy.Context

func ExampleDeepEqualTo() {
	is := inspect.With(t)

	is(0, DeepEqualTo(1))
	// Output:
	// provided: 0
	// expected: 1
}

func ExampleEqualTo_true() {
	is := inspect.With(t)

	is(false, EqualTo(true))
	// Output:
	// provided: false
	// expected: true
}

func ExampleEqualTo_one() {
	is := inspect.With(t)

	is(0, EqualTo(1))
	// Output:
	// provided: 0
	// expected: 1
}

func ExampleEqualTo_struct() {
	is := inspect.With(t)

	a := struct{ int }{0}
	b := struct{ int64 }{0}

	is(a, EqualTo(b))
	// Output:
	// provided: struct { int }{int:0}
	// expected: struct { int64 }{int64:0}
}

func ExampleEqualTo_string() {
	is := inspect.With(t)

	is("something", EqualTo("else"))
	// Output:
	// provided: "something"
	// expected: "else"
}

func ExampleError() {
	is := inspect.With(t)

	is(io.ErrUnexpectedEOF, Error(io.EOF))

	// Output:
	// provided: &errors.errorString{s:"unexpected EOF"}
	// expected: &errors.errorString{s:"EOF"}
}

func ExampleInPanic() {
	is := inspect.With(t)

	is(func() {
		// missing panic
	}, InPanic())

	// Output:
	// provided: github.com/4rcode/gadget/really_test.ExampleInPanic.func1
	// expected: panic
}

func ExampleInPanic_usePanicValue() {
	is := inspect.With(t)

	var v interface{}

	is(func() {
		panic("xyz")
	}, InPanic(&v))

	is(v, "abc")

	// Output:
	// provided: "xyz"
	// expected: "abc"
}

func ExampleNil() {
	is := inspect.With(t)

	is(123, Nil())

	// Output:
	// provided: 123
	// expected: <nil>
}

func ExampleNot_nil() {
	is := inspect.With(t)

	is((*struct{})(nil), Not(Nil()))
	// Output:
	// provided: (*struct {})(nil)
	// expected: not <nil>
}

func ExampleTrue() {
	is := inspect.With(t)

	is(0 >= 1, True("%v <=> %v", 0, 1))
	// Output:
	// provided: false
	// expected: 0 <=> 1
}
