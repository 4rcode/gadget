package mock_test

import (
	"fmt"

	"github.com/4rcode/gadget/mock"
)

func ExampleReporter() {
	var r mock.Reporter
	var detector mock.Detector = "detector"

	r.Helper()
	r.Report("=>", detector.Detect(123, nil), nil)

	fmt.Println(r)
	// Output:
	// #=> detector
}
