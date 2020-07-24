package pleat

import (
	"context"
	"io"
	"time"
)

var (
	// std is the name of the standard logger in stdlib `log`
	std = New()
)

func StandardLogger() *Logger {
	return std
}

// SetOutput ...
func SetOutput(out io.Writer) {
	std.SetOutput(out)
}

// SetFormatter ...
func SetFormatter(formatter Formatter) {
	std.SetFormatter(formatter)
}

// SetReportCaller ...
func SetReportCaller(include bool) {
	std.SetReportCaller(include)
}

// SetLevel ...
func SetLevel(level Level) {
	std.SetLevel(level)
}

// GetLevel ...
func GetLevel() Level {
	panic("GetLevel not implemented...yet...")
	return InfoLevel
}

// IsLevelEnabled ...
func IsLevelEnabled(level Level) bool {
	panic("IsLevelEnabled not implemented...yet...")
	return false
}

// AddHook ...
func AddHook(hook Hook) {
	panic("AddHook not implemented...yet...")
}

// WithError ...
func WithError(err error) *Entry {
	panic("WithError not implemented...yet...")
}

// WithContext ...
func WithContext(ctx context.Context) *Entry {
	panic("WithContext not implemented...yet...")
}

// WithField ...
func WithField(key string, value interface{}) *Entry {
	panic("WithField not implemented...yet...")
}

// WithFields ...
func WithFields(fields Fields) *Entry {
	panic("WithFields not implemented...yet...")
}

// WithTime ...
func WithTime(t time.Time) *Entry {
	panic("WithTime not implemented...yet...")
}

// Trace ...
func Trace(args ...interface{}) {
	panic("Trace not implemented...yet...")
}

// Debug ...
func Debug(args ...interface{}) {
	panic("Debug not implemented...yet...")
}

// Print ...
func Print(args ...interface{}) {
	panic("Print not implemented...yet...")
}

// Info ...
func Info(args ...interface{}) {
	panic("Info not implemented...yet...")
}

// Warn ...
func Warn(args ...interface{}) {
	panic("Warn not implemented...yet...")
}

// Warning ...
func Warning(args ...interface{}) {
	panic("Warning not implemented...yet...")
}

// Error ...
func Error(args ...interface{}) {
	panic("Error not implemented...yet...")
}

// Panic ...
func Panic(args ...interface{}) {
	panic("Panic not implemented...yet...")
}

// Fatal ...
func Fatal(args ...interface{}) {
	panic("Fatal not implemented...yet...")
}

// TraceFn ...
func TraceFn(fn LogFunction) {
	panic("TraceFn not implemented...yet...")
}

// DebugFn ...
func DebugFn(fn LogFunction) {
	panic("DebugFn not implemented...yet...")
}

// PrintFn ...
func PrintFn(fn LogFunction) {
	panic("PrintFn not implemented...yet...")
}

// InfoFn ...
func InfoFn(fn LogFunction) {
	panic("InfoFn not implemented...yet...")
}

// WarnFn ...
func WarnFn(fn LogFunction) {
	panic("WarnFn not implemented...yet...")
}

// WarningFn ...
func WarningFn(fn LogFunction) {
	panic("WarningFn not implemented...yet...")
}

// ErrorFn ...
func ErrorFn(fn LogFunction) {
	panic("ErrorFn not implemented...yet...")
}

// PanicFn ...
func PanicFn(fn LogFunction) {
	panic("PanicFn not implemented...yet...")
}

// FatalFn ...
func FatalFn(fn LogFunction) {
	panic("FatalFn not implemented...yet...")
}

// Tracef ...
func Tracef(format string, args ...interface{}) {
	panic("Tracef not implemented...yet...")
}

// Debugf ...
func Debugf(format string, args ...interface{}) {
	panic("Debugf not implemented...yet...")
}

// Printf ...
func Printf(format string, args ...interface{}) {
	panic("Printf not implemented...yet...")
}

// Infof ...
func Infof(format string, args ...interface{}) {
	panic("Infof not implemented...yet...")
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	panic("Warnf not implemented...yet...")
}

// Warningf ...
func Warningf(format string, args ...interface{}) {
	panic("Warningf not implemented...yet...")
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	panic("Errorf not implemented...yet...")
}

// Panicf ...
func Panicf(format string, args ...interface{}) {
	panic("Panicf not implemented...yet...")
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	panic("Fatalf not implemented...yet...")
}

// Traceln ...
func Traceln(args ...interface{}) {
	panic("Traceln not implemented...yet...")
}

// Debugln ...
func Debugln(args ...interface{}) {
	panic("Debugln not implemented...yet...")
}

// Println ...
func Println(args ...interface{}) {
	panic("Println not implemented...yet...")
}

// Infoln ...
func Infoln(args ...interface{}) {
	panic("Infoln not implemented...yet...")
}

// Warnln ...
func Warnln(args ...interface{}) {
	panic("Warnln not implemented...yet...")
}

// Warningln ...
func Warningln(args ...interface{}) {
	panic("Warningln not implemented...yet...")
}

// Errorln ...
func Errorln(args ...interface{}) {
	panic("Errorln not implemented...yet...")
}

// Panicln ...
func Panicln(args ...interface{}) {
	panic("Panicln not implemented...yet...")
}

// Fatalln ...
func Fatalln(args ...interface{}) {
	panic("Fatalln not implemented...yet...")
}
