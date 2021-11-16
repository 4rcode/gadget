package really_test

import (
	"testing"

	"github.com/4rcode/gadget/inspect"
	. "github.com/4rcode/gadget/really"
)

func TestInPanic(t *testing.T) {
	is := inspect.With(t)

	is(InPanic().Detect(nil, nil), String("func()"))
	is(InPanic().Detect(123, nil), String("func()"))
	is(InPanic().Detect(true, nil), String("func()"))
	is(InPanic().Detect((func())(nil), nil), String("func()"))

	var a, b, c interface{}

	detector := InPanic(nil, &a, &b, &c)

	for testCaseId, expected := range []interface{}{
		(*int)(nil), 0, "err", nil,
	} {
		t.Log(testCaseId + 1)

		err := detector.Detect(func() {
			if expected != nil {
				panic(expected)
			}
		}, nil)

		if expected == nil {
			is(err, String("panic"))
		}

		is(a, expected)
		is(b, expected)
		is(c, expected)
	}
}
