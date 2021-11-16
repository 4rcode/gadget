package mock_test

import (
	"fmt"

	"github.com/4rcode/gadget/mock"
)

func ExampleBuilder() {
	var b mock.Builder

	fmt.Println(b.Build(1, 2, 3))

	b = "=> "

	fmt.Println(b.Build("a", "b", "c"))
	// Output:
	// 1 2 3
	// => abc
}
