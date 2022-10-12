package iox

import (
	"bytes"
	"io"
)

var _ Stderrer = &Buffer{}
var _ Stdiner = &Buffer{}
var _ Stdioer = &Buffer{}
var _ Stdouter = &Buffer{}

type Buffer struct {
	In  io.Reader
	Out bytes.Buffer
	Err bytes.Buffer
}

func (oi *Buffer) Stdout() io.Writer {
	return &oi.Out
}

func (oi *Buffer) Stderr() io.Writer {
	return &oi.Err
}

func (oi *Buffer) Stdin() io.Reader {
	return oi.In
}

func (oi Buffer) String() string {
	return IO{
		In:  oi.Stdin(),
		Out: oi.Stdout(),
		Err: oi.Stderr(),
	}.String()
}

func (oi Buffer) MarshalJSON() ([]byte, error) {
	return IO{
		In:  oi.Stdin(),
		Out: oi.Stdout(),
		Err: oi.Stderr(),
	}.MarshalJSON()
}
