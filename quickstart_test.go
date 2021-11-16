package gadget_test

import (
	"github.com/4rcode/gadget/inspect"
	. "github.com/4rcode/gadget/really"
)

func Example() {

	//
	// Initialize an inspector in your test.
	//

	is := inspect.With(t)

	//
	// GIVEN your test subject and its dependencies.
	//

	var (
		err    error
		result int
	)

	//
	// WHEN you have executed your logic.
	//

	result = 1 + 1

	//
	// THEN you can make assertions.
	//

	is(result == 2)
	is(result == 2, True("equal to two"))
	is(result, 2)
	is(result, 2.0)
	is(result, String("2"))
	is(result, Not(Nil()))
	is(result, Not(Error(err)))
	is(result, Not(InPanic()))
	is(result, EqualTo(2, 2.0))
	is(result, DeepEqualTo(2))
	is(result, Like(float64(0)))
	is(result, Either(1, 2, 3), Not(4, 5, 6), EitherNot(2, 2.0, 0))

	is(err, Zero())

	// Output:

}
