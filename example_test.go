package gadget_test

import (
	"math/rand"

	"github.com/4rcode/gadget/dummy"
	"github.com/4rcode/gadget/inspect"
	. "github.com/4rcode/gadget/really"
)

var t dummy.Context

// The first argument can simply be a boolean expression. When no other
// arguments are supplied, the default behavior is to detect a true value.
func Example_a() {
	is := inspect.With(t)

	is(rand.Intn(5) > 10)
	// Output:
	// provided: false
	// expected: true
}

// Optionally, you can use the True Detector to specify a custom message.
func Example_b() {
	is := inspect.With(t)

	is(rand.Intn(5) > 10, True("a value greater than %v", 10))
	// Output:
	// provided: false
	// expected: a value greater than 10
}

// Based on the natural way of reading, the first argument is the tested
// subject, followed by one or more target objects.
//
// Values will be compared for equality by default.
func Example_c() {
	is := inspect.With(t)

	is(rand.Intn(1), 1)
	// Output:
	// provided: 0
	// expected: 1
}

// There are many Detectors you can use, just check out the "really" package. We
// suggest to import the entire set of Detectors using the DOT notation:
//
//   import . "github.com/4rcode/gadget/really"
func Example_d() {
	is := inspect.With(t)

	x := new(struct{ int })
	y := new(struct{ int })

	is(x, EqualTo(y))     // it will fail & report
	is(x, DeepEqualTo(y)) // it is true!
	// Output:
	// provided: &struct { int }{int:0}
	// expected: &struct { int }{int:0}
}

// Do not forget that you can specify more than one Detector at once!
func Example_e() {
	is := inspect.With(t)

	is(123, Not(Nil()), Like(true))
	// Output:
	// provided: 123
	// expected: bool instance
}
