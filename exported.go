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
	panic("not implemented...yet...")
	return InfoLevel
}

// IsLevelEnabled ...
func IsLevelEnabled(level Level) bool {
	panic("not implemented...yet...")
	return false
}

// AddHook ...
func AddHook(hook Hook) {
	panic("not implemented...yet...")
}

// WithError ...
func WithError(err error) *Entry {
	panic("not implemented...yet...")
}

// WithContext ...
func WithContext(ctx context.Context) *Entry {
	panic("not implemented...yet...")
}

// WithField ...
func WithField(key string, value interface{}) *Entry {
	panic("not implemented...yet...")
}

// WithFields ...
func WithFields(fields Fields) *Entry {
	panic("not implemented...yet...")
}

// WithTime ...
func WithTime(t time.Time) *Entry {
	panic("not implemented...yet...")
}

// Trace ...
func Trace(args ...interface{}) {
	panic("not implemented...yet...")
}

// Debug ...
func Debug(args ...interface{}) {
	panic("not implemented...yet...")
}

// Print ...
func Print(args ...interface{}) {
	panic("not implemented...yet...")
}

// Info ...
func Info(args ...interface{}) {
	panic("not implemented...yet...")
}

// Warn ...
func Warn(args ...interface{}) {
	panic("not implemented...yet...")
}

// Warning ...
func Warning(args ...interface{}) {
	panic("not implemented...yet...")
}

// Error ...
func Error(args ...interface{}) {
	panic("not implemented...yet...")
}

// Panic ...
func Panic(args ...interface{}) {
	panic("not implemented...yet...")
}

// Fatal ...
func Fatal(args ...interface{}) {
	panic("not implemented...yet...")
}

// TraceFn ...
func TraceFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// DebugFn ...
func DebugFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// PrintFn ...
func PrintFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// InfoFn ...
func InfoFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// WarnFn ...
func WarnFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// WarningFn ...
func WarningFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// ErrorFn ...
func ErrorFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// PanicFn ...
func PanicFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// FatalFn ...
func FatalFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// Tracef ...
func Tracef(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Debugf ...
func Debugf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Printf ...
func Printf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Infof ...
func Infof(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Warningf ...
func Warningf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Panicf ...
func Panicf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Traceln ...
func Traceln(args ...interface{}) {
	panic("not implemented...yet...")
}

// Debugln ...
func Debugln(args ...interface{}) {
	panic("not implemented...yet...")
}

// Println ...
func Println(args ...interface{}) {
	panic("not implemented...yet...")
}

// Infoln ...
func Infoln(args ...interface{}) {
	panic("not implemented...yet...")
}

// Warnln ...
func Warnln(args ...interface{}) {
	panic("not implemented...yet...")
}

// Warningln ...
func Warningln(args ...interface{}) {
	panic("not implemented...yet...")
}

// Errorln ...
func Errorln(args ...interface{}) {
	panic("not implemented...yet...")
}

// Panicln ...
func Panicln(args ...interface{}) {
	panic("not implemented...yet...")
}

// Fatalln ...
func Fatalln(args ...interface{}) {
	panic("not implemented...yet...")
}
