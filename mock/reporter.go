package mock

import (
	"fmt"

	"github.com/4rcode/gadget"
)

type Reporter string

func (r *Reporter) Helper() {
	*r = *r + "#"
}

func (r *Reporter) Report(value interface{}, err error, options gadget.Options) {
	*r = *r + Reporter(fmt.Sprint(value, " ", err))
}
