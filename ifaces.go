package iox

import "io"

// IOable can be implemented to return an IO.
type IOable interface {
	Stdio() IO
}

// IOSetable can be implemented to receive an IO.
type IOSetable interface {
	SetStdio(oi IO) error
}

type Stdiner interface {
	Stdin() io.Reader
}

type Stdouter interface {
	Stdout() io.Writer
}

type Stderrer interface {
	Stderr() io.Writer
}

type Stdioer interface {
	Stderrer
	Stdiner
	Stdouter
}
