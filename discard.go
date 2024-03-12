package iox

import (
	"bytes"
	"io"
)

// Discard returns an IO
// that discards Stderr
// and Stdout.
func Discard() IO {
	return IO{
		Err: io.Discard,
		In:  bytes.NewReader(nil),
		Out: io.Discard,
	}
}
