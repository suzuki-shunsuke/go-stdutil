package stdutil

import (
	"io"
)

type LDFlags struct {
	Version string
	Commit  string
	Date    string
}

type Env struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
	Getenv Getenv
}
