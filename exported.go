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
	panic("not implemented...yet...")
	return nil
}

// SetOutput sets the standard logger output.
func SetOutput(out io.Writer) {
	panic("not implemented...yet...")
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter Formatter) {
	panic("not implemented...yet...")
}

// SetReportCaller sets whether the standard logger will include the calling
// method as a field.
func SetReportCaller(include bool) {
	panic("not implemented...yet...")
}

// SetLevel sets the standard logger level.
func SetLevel(level Level) {
	panic("not implemented...yet...")
}

// GetLevel returns the standard logger level.
func GetLevel() Level {
	panic("not implemented...yet...")
	return InfoLevel
}

// IsLevelEnabled checks if the log level of the standard logger is greater than the level param
func IsLevelEnabled(level Level) bool {
	panic("not implemented...yet...")
	return false
}

// AddHook adds a hook to the standard logger hooks.
func AddHook(hook Hook) {
	panic("not implemented...yet...")
}

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) *Entry {
	panic("not implemented...yet...")
}

// WithContext creates an entry from the standard logger and adds a context to it.
func WithContext(ctx context.Context) *Entry {
	panic("not implemented...yet...")
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *Entry {
	panic("not implemented...yet...")
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields Fields) *Entry {
	panic("not implemented...yet...")
}

// WithTime creates an entry from the standard logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithTime(t time.Time) *Entry {
	panic("not implemented...yet...")
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	panic("not implemented...yet...")
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	panic("not implemented...yet...")
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	panic("not implemented...yet...")
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	panic("not implemented...yet...")
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	panic("not implemented...yet...")
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	panic("not implemented...yet...")
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	panic("not implemented...yet...")
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	panic("not implemented...yet...")
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	panic("not implemented...yet...")
}

// TraceFn logs a message from a func at level Trace on the standard logger.
func TraceFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// DebugFn logs a message from a func at level Debug on the standard logger.
func DebugFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// PrintFn logs a message from a func at level Info on the standard logger.
func PrintFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// InfoFn logs a message from a func at level Info on the standard logger.
func InfoFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// WarnFn logs a message from a func at level Warn on the standard logger.
func WarnFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// WarningFn logs a message from a func at level Warn on the standard logger.
func WarningFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// ErrorFn logs a message from a func at level Error on the standard logger.
func ErrorFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// PanicFn logs a message from a func at level Panic on the standard logger.
func PanicFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// FatalFn logs a message from a func at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalFn(fn LogFunction) {
	panic("not implemented...yet...")
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

// Traceln logs a message at level Trace on the standard logger.
func Traceln(args ...interface{}) {
	panic("not implemented...yet...")
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	panic("not implemented...yet...")
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	panic("not implemented...yet...")
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	panic("not implemented...yet...")
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	panic("not implemented...yet...")
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
	panic("not implemented...yet...")
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	panic("not implemented...yet...")
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	panic("not implemented...yet...")
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(args ...interface{}) {
	panic("not implemented...yet...")
}
