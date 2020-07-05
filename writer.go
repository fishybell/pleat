package pleat

import (
	"bufio"
	"io"
	"runtime"
)

// Writer ...
func (logger *Logger) Writer() *io.PipeWriter {
	return logger.WriterLevel(logger.Level)
}

// WriterLevel ...
func (logger *Logger) WriterLevel(level Level) *io.PipeWriter {
	return NewEntry(logger).WriterLevel(level)
}

// Writer ...
func (entry *Entry) Writer() *io.PipeWriter {
	return entry.WriterLevel(entry.Level)
	return nil
}

// WriterLevel ...
func (entry *Entry) WriterLevel(level Level) *io.PipeWriter {
	reader, writer := io.Pipe()

	runtime.SetFinalizer(writer, writerFinalizer)
	go entry.pipeReader(reader, level)

	return writer
}

func (entry *Entry) pipeReader(reader *io.PipeReader, level Level) {
	scanner := bufio.NewScanner(reader)

	var print func(args ...interface{})

	switch level {
	case PanicLevel:
		print = entry.Panic
	case FatalLevel:
		print = entry.Fatal
	case ErrorLevel:
		print = entry.Error
	case WarnLevel:
		print = entry.Warn
	case InfoLevel:
		print = entry.Info
	case DebugLevel:
		print = entry.Debug
	case TraceLevel:
		print = entry.Trace
	}

	for scanner.Scan() {
		print(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		entry.WithError(err).Error("Pleat pipe failed to read")
	}

	reader.Close()
}

func writerFinalizer(writer *io.PipeWriter) {
	writer.Close()
}
