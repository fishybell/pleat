package pleat

import (
	"log"

	"github.com/go-errors/errors"
)

// Fields ...
type Fields map[string]interface{}

// Level ...
type Level uint32

// String ...
func (level Level) String() string {
	bytes, _ := level.MarshalText()
	return string(bytes)
}

// ParseLeve ...
func ParseLevel(lvl string) (Level, error) {
	panic("string.( not implemented...yet...")
	return InfoLevel, nil
}

// UnmarshalText ...
func (level *Level) UnmarshalText(text []byte) error {
	panic("Level.UnmarshalText not implemented...yet...")
	return nil
}

func (level Level) MarshalText() ([]byte, error) {
	switch level {
	case PanicLevel:
		return []byte("panic"), nil
	case FatalLevel:
		return []byte("fatal"), nil
	case ErrorLevel:
		return []byte("error"), nil
	case WarnLevel:
		return []byte("warn"), nil
	case InfoLevel:
		return []byte("info"), nil
	case DebugLevel:
		return []byte("debug"), nil
	case TraceLevel:
		return []byte("trace"), nil
	}

	return []byte("unknown"), errors.Errorf("Invalid Log Level")
}

// AllLevels ...
var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
	TraceLevel,
}

// Levels ...
const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

// Won't compile if StdLogger can't be realized by a log.Logger
var (
	_ StdLogger = &log.Logger{}
	_ StdLogger = &Entry{}
	_ StdLogger = &Logger{}
)

// StdLogger ...
type StdLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}

// FieldLogger ...
type FieldLogger interface {
	WithField(key string, value interface{}) *Entry
	WithFields(fields ...interface{}) *Entry
	WithError(err error) *Entry

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})
}
