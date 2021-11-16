package inspect_test

import (
	"math/rand"

	"github.com/4rcode/gadget/dummy"
	"github.com/4rcode/gadget/inspect"
	. "github.com/4rcode/gadget/really"
	"github.com/4rcode/gadget/settings"
)

var t dummy.Context

func ExampleWith() {
	is := inspect.With(t)

	is(rand.Intn(1) == 1)
	// Output:
	// provided: false
	// expected: true
}

func ExampleWith_customTemplate() {
	is := inspect.With(t, settings.Template("%[1]v <> %[2]v"))

	is(rand.Intn(1) == 1)
	// Output:
	// false <> true
}

func ExampleWithExclusive() {
	is := inspect.WithExclusive(t)

	is(0, 1)
	// Output:
	// provided: 0
	// expected: 1
}

func ExampleWithLenient() {
	is := inspect.WithLenient(t)

	is(nil, Not(Nil()))
	// Output:
	// provided: <nil>
	// expected: not <nil>
}

func ExampleWithStrict() {
	is := inspect.WithStrict(t)

	is("abc", "xyz")
	// Output:
	// provided: "abc"
	// expected: "xyz"
}
