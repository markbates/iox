package iox

import "io"

// Discard returns an IO
// that discards sends Stderr
// and Stdout to io.Discard.
func Discard() IO {
	return IO{
		Out: io.Discard,
		Err: io.Discard,
	}
}
