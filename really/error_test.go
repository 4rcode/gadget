package really_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/4rcode/gadget/inspect"
	. "github.com/4rcode/gadget/really"
	"github.com/4rcode/gnomic"
)

func TestError(t *testing.T) {
	is := inspect.With(t)

	type err struct{ error }

	custom := fmt.Errorf("error %w", io.EOF)

	for ti, td := range []struct {
		in interface{}
		error
		out bool
	}{
		{nil, nil, true},
		{nil, io.EOF, false},
		{nil, custom, false},
		{nil, (*err)(nil), false},
		{error(nil), nil, true},
		{error(nil), io.EOF, false},
		{error(nil), custom, false},
		{error(nil), (*err)(nil), false},
		{(*err)(nil), nil, false},
		{(*err)(nil), io.EOF, false},
		{(*err)(nil), custom, false},
		{(*err)(nil), (*err)(nil), true},
		{io.EOF, nil, false},
		{io.EOF, io.EOF, true},
		{io.EOF, custom, false},
		{io.EOF, (*err)(nil), false},
		{io.EOF, io.ErrClosedPipe, false},
		{custom, nil, false},
		{custom, io.EOF, true},
		{custom, custom, true},
		{custom, (*err)(nil), false},
		{custom, io.ErrClosedPipe, false},
	} {
		t.Log(ti+1, td)

		var opts gnomic.Options

		is(Error(td.error).Detect(td.in, opts) == nil, td.out)
	}
}
