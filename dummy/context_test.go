package dummy_test

import (
	"github.com/4rcode/gadget/dummy"
)

func ExampleContext() {
	var c dummy.Context

	c.Helper()
	c.Errorf("%s ", "a")
	c.Fatalf("%d ", 1)
	c.Logf("%.2f ", 2.3)
	c.Skipf("%#v", struct{}{})
	// Output:
	// a 1 2.30 struct {}{}
}
