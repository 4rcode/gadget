# gadget - assertion library

[![
	 Build
](https://github.com/4rcode/gadget/actions/workflows/build.yml/badge.svg)
](https://github.com/4rcode/gadget/actions/workflows/build.yml)
[![
	Coverage
](https://codecov.io/gh/4rcode/gadget/branch/main/graph/badge.svg)
](https://codecov.io/gh/4rcode/gadget/branch/main)
[![
	Reference
](https://pkg.go.dev/badge/github.com/4rcode/gadget.svg)
](https://pkg.go.dev/github.com/4rcode/gadget)

## Installation

```sh
go get github.com/4rcode/gadget
```

## [Quick Start](quickstart_test.go)

```go
import (
	"testing"

	"github.com/4rcode/gadget/inspect"
	. "github.com/4rcode/gadget/really"
)

func TestWithGadget(t *testing.T) {
	is := inspect.With(t)

	var err error

	result := 1 + 1

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
}
```
