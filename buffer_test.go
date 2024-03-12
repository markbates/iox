package iox

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Buffer(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var nb *Buffer
	r.Equal(IO{}, nb.IO())

	buf := Buffer{
		In: strings.NewReader("hello"),
	}

	oi := buf.IO()

	fmt.Fprint(oi.Stdout(), "STDOUT")
	fmt.Fprint(oi.Stderr(), "STDERR")

	r.Equal("STDOUT", buf.Out.String())
	r.Equal("STDERR", buf.Err.String())

	b, err := io.ReadAll(oi.Stdin())
	r.NoError(err)
	r.Equal("hello", string(b))
}
