package iox

import (
	"bytes"
	"io"
)

// Buffer is useful for testing.
type Buffer struct {
	In  io.Reader
	Out bytes.Buffer
	Err bytes.Buffer
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

func (oi Buffer) MarshalJSON() ([]byte, error) {
	return oi.IO().MarshalJSON()
}

func (oi Buffer) String() string {
	return oi.IO().String()
}
