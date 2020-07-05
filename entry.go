package pleat

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"time"
)

// ErrorKey ...
var ErrorKey = "error"

// Entry ...
type Entry struct {
	Logger  *Logger
	Data    Fields
	Time    time.Time
	Level   Level
	Caller  *runtime.Frame
	Message string
	Buffer  *bytes.Buffer
	Context context.Context
}

// NewEntry ...
func NewEntry(logger *Logger) *Entry {
	return &Entry{
		Logger: logger,
	}
}

// Bytes ...
func (entry *Entry) Bytes() ([]byte, error) {
	panic("not implemented...yet...")
	return nil, nil
}

// String ...
func (entry *Entry) String() (string, error) {
	panic("not implemented...yet...")
	return "", nil
}

// WithError ...
func (entry *Entry) WithError(err error) *Entry {
	panic("not implemented...yet...")
	return nil
}

// WithContext ...
func (entry *Entry) WithContext(ctx context.Context) *Entry {
	panic("not implemented...yet...")
	return nil
}

// WithField ...
func (entry *Entry) WithField(key string, value interface{}) *Entry {
	panic("not implemented...yet...")
	return nil
}

// WithFields ...
func (entry *Entry) WithFields(fields ...interface{}) *Entry {
	newEntry := *entry

	for _, fieldUnknown := range fields {
		switch field := fieldUnknown.(type) {
		case Fields:
			newEntry.Data = appendData(entry.Data, field)
		case map[string]interface{}:
			newEntry.Data = appendData(entry.Data, Fields(field))
		default:
			panic("not implemented...yet...")
		}
	}

	return &newEntry
}

func appendData(original, changes Fields) Fields {
	for k, v := range changes {
		original[k] = v
	}

	return original
}

// WithTime ...
func (entry *Entry) WithTime(t time.Time) *Entry {
	panic("not implemented...yet...")
	return nil
}

func (entry Entry) HasCaller() (has bool) {
	panic("not implemented...yet...")
	return
}

func (entry Entry) log(level Level, args ...interface{}) {
	// non-pointer function so changees don't effect future entries
	entry.Level = level
	entry.Message = fmt.Sprint(args...)
	entry.write()
}

func (entry *Entry) write() {
	entry.Logger.mu.Lock()
	defer entry.Logger.mu.Unlock()

	formatted, err := entry.Logger.Formatter.Format(entry)
	if err != nil {
		fmt.Fprintf(entry.Logger.ErrorOut, "Formatter failed to format, %v\n", err)
		return
	}

	if _, err = entry.Logger.Out.Write(formatted); err != nil {
		fmt.Fprintf(entry.Logger.ErrorOut, "Pleat failed to write log, %v\n", err)
	}
}

// Log ...
func (entry *Entry) Log(level Level, args ...interface{}) {
	if entry.Logger.IsLevelEnabled(level) {
		entry.log(level, args...)
	}
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
	entry.Log(InfoLevel, args...)
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
