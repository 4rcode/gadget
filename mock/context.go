package mock

import (
	"fmt"
)

type Context string

func (c *Context) Errorf(f string, a ...interface{}) {
	c.Logf(f, a...)
}

func (c *Context) Fatalf(f string, a ...interface{}) {
	c.Logf(f, a...)
}

func (c *Context) Helper() {
	c.Logf("#")
}

func (c *Context) Logf(f string, a ...interface{}) {
	*c = *c + Context(fmt.Sprintf(f, a...))
}

func (c *Context) Skipf(f string, a ...interface{}) {
	c.Logf(f, a...)
}

func (c Context) String() string {
	return string(c)
}
