package dummy

import (
	"fmt"
)

type Context struct{}

func (c Context) Errorf(format string, arguments ...interface{}) {
	c.Logf(format, arguments...)
}

func (c Context) Fatalf(format string, arguments ...interface{}) {
	c.Logf(format, arguments...)
}

func (Context) Helper() {}

func (c Context) Logf(format string, arguments ...interface{}) {
	fmt.Printf(format, arguments...)
}

func (c Context) Skipf(format string, arguments ...interface{}) {
	c.Logf(format, arguments...)
}
