package iox

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var _ Stderrer = IO{}
var _ Stdiner = IO{}
var _ Stdioer = IO{}
var _ Stdouter = IO{}

// IO represents the standard input, output, and error stream.
type IO struct {
	In  io.Reader     `json:"in,omitempty"`  // standard input
	Out io.ReadWriter `json:"out,omitempty"` // standard output
	Err io.ReadWriter `json:"err,omitempty"` // standard error
}

// Stdout returns IO.In.
// Defaults to os.Stdout.
func (oi IO) Stdout() io.ReadWriter {
	if oi.Out == nil {
		return os.Stdout
	}

	return oi.Out
}

// Stderr returns IO.Err.
// Defaults to os.Stderr.
func (oi IO) Stderr() io.ReadWriter {
	if oi.Err == nil {
		return os.Stderr
	}

	return oi.Err
}

// Stdin returns IO.In.
// Defaults to os.Stdin.
func (oi IO) Stdin() io.Reader {
	if oi.In == nil {
		return os.Stdin
	}

	return oi.In
}

func (oi IO) String() string {
	b, _ := json.MarshalIndent(oi, "", "  ")
	return string(b)
}

func (oi IO) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		"stdin":  fmt.Sprintf("%T", oi.Stdin()),
		"stdout": fmt.Sprintf("%T", oi.Stdout()),
		"stderr": fmt.Sprintf("%T", oi.Stderr()),
	}

	return json.Marshal(m)
}
