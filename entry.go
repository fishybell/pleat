package pleat

import (
	"bytes"
	"context"
	"runtime"
	"time"
)

// Defines the key when adding errors using WithError.
var ErrorKey = "error"

// An entry is the final or intermediate Logrus logging entry. It contains all
// the fields passed with WithField{,s}. It's finally logged when Trace, Debug,
// Info, Warn, Error, Fatal or Panic is called on it. These objects can be
// reused and passed around as much as you wish to avoid field duplication.
type Entry struct {
	Logger *Logger

	// Contains all the fields set by the user.
	Data Fields

	// Time at which the log entry was created
	Time time.Time

	// Level the log entry was logged at: Trace, Debug, Info, Warn, Error, Fatal or Panic
	// This field will be set on entry firing and the value will be equal to the one in Logger struct field.
	Level Level

	// Calling method, with package name
	Caller *runtime.Frame

	// Message passed to Trace, Debug, Info, Warn, Error, Fatal or Panic
	Message string

	// When formatter is called in entry.log(), a Buffer may be set to entry
	Buffer *bytes.Buffer

	// Contains the context set by the user. Useful for hook processing etc.
	Context context.Context

	// err may contain a field formatting error
	err string
}

func NewEntry(logger *Logger) *Entry {
	panic("not implemented...yet...")
	return nil
}

// Returns the bytes representation of this entry from the formatter.
func (entry *Entry) Bytes() ([]byte, error) {
	panic("not implemented...yet...")
	return nil, nil
}

// Returns the string representation from the reader and ultimately the
// formatter.
func (entry *Entry) String() (string, error) {
	panic("not implemented...yet...")
	return "", nil
}

// Add an error as single field (using the key defined in ErrorKey) to the Entry.
func (entry *Entry) WithError(err error) *Entry {
	panic("not implemented...yet...")
	return nil
}

// Add a context to the Entry.
func (entry *Entry) WithContext(ctx context.Context) *Entry {
	panic("not implemented...yet...")
	return nil
}

// Add a single field to the Entry.
func (entry *Entry) WithField(key string, value interface{}) *Entry {
	panic("not implemented...yet...")
	return nil
}

// Add a map of fields to the Entry.
func (entry *Entry) WithFields(fields Fields) *Entry {
	panic("not implemented...yet...")
	return nil
}

// Overrides the time of the Entry.
func (entry *Entry) WithTime(t time.Time) *Entry {
	panic("not implemented...yet...")
	return nil
}

func (entry Entry) HasCaller() (has bool) {
	panic("not implemented...yet...")
	return
}

// This function is not declared with a pointer value because otherwise
// race conditions will occur when using multiple goroutines
func (entry Entry) log(level Level, msg string) {
	panic("not implemented...yet...")
}

func (entry *Entry) Log(level Level, args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Trace(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Debug(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Print(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Info(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Warn(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Warning(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Error(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Fatal(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Panic(args ...interface{}) {
	panic("not implemented...yet...")
}

// Entry Printf family functions

func (entry *Entry) Logf(level Level, format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Tracef(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Debugf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Infof(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Printf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Warnf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Warningf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Errorf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Fatalf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Panicf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Entry Println family functions

func (entry *Entry) Logln(level Level, args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Traceln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Debugln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Infoln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Println(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Warnln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Warningln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Errorln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Fatalln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (entry *Entry) Panicln(args ...interface{}) {
	panic("not implemented...yet...")
}
