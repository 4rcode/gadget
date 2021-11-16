package mock_test

import (
	"fmt"

	"github.com/4rcode/gadget/mock"
)

func ExampleContext() {
	var c mock.Context

	c.Helper()
	c.Errorf(" %s", "a")
	c.Fatalf(" %d", 1)
	c.Logf(" %.2f", 2.3)
	c.Skipf(" %#v", struct{}{})

	fmt.Println(c)
	// Output:
	// # a 1 2.30 struct {}{}
}
