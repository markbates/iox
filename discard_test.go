package iox

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Discard(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	d := Discard()
	r.NotNil(d.Out)
	r.NotNil(d.Err)
	r.NotNil(d.In)

	r.Equal(io.Discard, d.Out)
	r.Equal(io.Discard, d.Err)
}
