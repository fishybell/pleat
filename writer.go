package pleat

import (
	"io"
)

// Writer at INFO level. See WriterLevel for details.
func (logger *Logger) Writer() *io.PipeWriter {
	panic("not implemented...yet...")
	return nil
}

// WriterLevel returns an io.Writer that can be used to write arbitrary text to
// the logger at the given log level. Each line written to the writer will be
// printed in the usual way using formatters and hooks. The writer is part of an
// io.Pipe and it is the callers responsibility to close the writer when done.
// This can be used to override the standard library logger easily.
func (logger *Logger) WriterLevel(level Level) *io.PipeWriter {
	panic("not implemented...yet...")
	return nil
}

func (entry *Entry) Writer() *io.PipeWriter {
	panic("not implemented...yet...")
	return nil
}

func (entry *Entry) WriterLevel(level Level) *io.PipeWriter {
	panic("not implemented...yet...")
	return nil
}
