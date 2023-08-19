package iox

import (
	"bytes"
	"io"
)

// Buffer is useful for testing.
type Buffer struct {
	In  io.Reader    `json:"-"`
	Out bytes.Buffer `json:"-"`
	Err bytes.Buffer `json:"-"`
}

func (oi *Buffer) IO() IO {
	if oi == nil {
		return IO{}
	}

	return IO{
		In:  oi.In,
		Out: &oi.Out,
		Err: &oi.Err,
	}
}
