package pleat

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/go-errors/errors"
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
		Data:   Fields{},
	}
}

// Bytes ...
func (entry *Entry) Bytes() ([]byte, error) {
	panic("Entry.Bytes not implemented...yet...")
	return nil, nil
}

// String ...
func (entry *Entry) String() (string, error) {
	panic("Entry.String not implemented...yet...")
	return "", nil
}

// WithError ...
func (entry *Entry) WithError(err error) *Entry {
	newEntry := *entry

	switch goError := err.(type) {
	case *errors.Error:
		// errors from the go-errors package have a stack trace built-in
		// use that trace
		entry.Data["error"] = goError.ErrorStack()
	default:
		entry.Data["error"] = err.Error()
	}

	return &newEntry
}

// WithContext ...
func (entry *Entry) WithContext(ctx context.Context) *Entry {
	panic("Entry.WithContext not implemented...yet...")
	return nil
}

// WithField ...
func (entry *Entry) WithField(key string, value interface{}) *Entry {
	panic("Entry.WithField not implemented...yet...")
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
			panic("default in Entry.WithFields not implemented...yet...")
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
	panic("Entry.WithTime not implemented...yet...")
	return nil
}

func (entry Entry) HasCaller() (has bool) {
	return entry.Caller != nil
}

func (entry Entry) log(level Level, args ...interface{}) {
	// non-pointer function so changees don't effect future entries
	entry.Level = level
	entry.Message = fmt.Sprint(args...)
	entry.Caller = getCaller()
	entry.write()
}

func (entry *Entry) write() {
	entry.Logger.mu.Lock()
	defer entry.Logger.mu.Unlock()

	if entry.Logger == nil || entry.Logger.Formatter == nil {
		fmt.Fprintf(entry.Logger.ErrorOut, "Formatter not initialized on Logger")
		return
	}

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
	entry.Log(TraceLevel, args...)
}

func (entry *Entry) Debug(args ...interface{}) {
	entry.Log(DebugLevel, args...)
}

func (entry *Entry) Print(args ...interface{}) {
	entry.Info(args...)
}

func (entry *Entry) Info(args ...interface{}) {
	entry.Log(InfoLevel, args...)
}

func (entry *Entry) Warn(args ...interface{}) {
	entry.Log(WarnLevel, args...)
}

func (entry *Entry) Warning(args ...interface{}) {
	entry.Log(WarnLevel, args...)
}

func (entry *Entry) Error(args ...interface{}) {
	entry.Log(ErrorLevel, args...)
}

func (entry *Entry) Fatal(args ...interface{}) {
	entry.Log(FatalLevel, args...)
}

func (entry *Entry) Panic(args ...interface{}) {
	entry.Log(PanicLevel, args...)
}

func (entry *Entry) Logf(level Level, format string, args ...interface{}) {
	panic("Entry.Logf not implemented...yet...")
}

func (entry *Entry) Tracef(format string, args ...interface{}) {
	panic("Entry.Tracef not implemented...yet...")
}

func (entry *Entry) Debugf(format string, args ...interface{}) {
	panic("Entry.Debugf not implemented...yet...")
}

func (entry *Entry) Infof(format string, args ...interface{}) {
	panic("Entry.Infof not implemented...yet...")
}

func (entry *Entry) Printf(format string, args ...interface{}) {
	panic("Entry.Printf not implemented...yet...")
}

func (entry *Entry) Warnf(format string, args ...interface{}) {
	panic("Entry.Warnf not implemented...yet...")
}

func (entry *Entry) Warningf(format string, args ...interface{}) {
	panic("Entry.Warningf not implemented...yet...")
}

func (entry *Entry) Errorf(format string, args ...interface{}) {
	panic("Entry.Errorf not implemented...yet...")
}

func (entry *Entry) Fatalf(format string, args ...interface{}) {
	panic("Entry.Fatalf not implemented...yet...")
}

func (entry *Entry) Panicf(format string, args ...interface{}) {
	panic("Entry.Panicf not implemented...yet...")
}

func (entry *Entry) Logln(level Level, args ...interface{}) {
	panic("Entry.Logln not implemented...yet...")
}

func (entry *Entry) Traceln(args ...interface{}) {
	panic("Entry.Traceln not implemented...yet...")
}

func (entry *Entry) Debugln(args ...interface{}) {
	panic("Entry.Debugln not implemented...yet...")
}

func (entry *Entry) Infoln(args ...interface{}) {
	panic("Entry.Infoln not implemented...yet...")
}

func (entry *Entry) Println(args ...interface{}) {
	panic("Entry.Println not implemented...yet...")
}

func (entry *Entry) Warnln(args ...interface{}) {
	panic("Entry.Warnln not implemented...yet...")
}

func (entry *Entry) Warningln(args ...interface{}) {
	panic("Entry.Warningln not implemented...yet...")
}

func (entry *Entry) Errorln(args ...interface{}) {
	panic("Entry.Errorln not implemented...yet...")
}

func (entry *Entry) Fatalln(args ...interface{}) {
	panic("Entry.Fatalln not implemented...yet...")
}

func (entry *Entry) Panicln(args ...interface{}) {
	panic("Entry.Panicln not implemented...yet...")
}

func getCaller() *runtime.Frame {
	// Ask runtime.Callers for up to 10 program counters, including runtime.Callers itself.
	pc := make([]uintptr, 20)
	n := runtime.Callers(0, pc)
	if n == 0 {
		// No program counters available. Stop now.
		// This can happen if the first argument to runtime.Callers is large.
		return nil
	}

	pc = pc[:n] // pass only valid program counters to runtime.CallersFrames
	frames := runtime.CallersFrames(pc)

	// Loop to get the frame we want
	// A fixed number of program counters can expand to an indefinite number of Frames.
	reachedPleatModule := false
	for {
		frame, more := frames.Next()
		if !more {
			break
		}

		if strings.Contains(frame.File, "pleat/") {
			reachedPleatModule = true
			continue
		} else if reachedPleatModule {
			// the first non-pleat frame is where the user called pleat from
			return &frame
		}

	}
	return nil
}
