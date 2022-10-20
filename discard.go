package iox

import (
	"fmt"
	"io"
)

// Discard returns an IO
// that discards Stderr
// and Stdout.
func Discard() IO {
	return IO{
		Out: discard{},
		Err: discard{},
	}
}

type discard struct{}

func (discard) Write(p []byte) (n int, err error) {
	return io.Discard.Write(p)
}

func (discard) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("cannot read from discard")
}
