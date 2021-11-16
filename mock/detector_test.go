package mock_test

import (
	"fmt"

	"github.com/4rcode/gadget/mock"
)

func ExampleDetector() {
	var d mock.Detector

	fmt.Println(d.Detect("", nil))

	d = "wrong"

	fmt.Println(d.Detect("", nil))
	// Output:
	// <nil>
	// wrong
}
