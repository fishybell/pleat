package pleat

import (
	"context"
	"io"
	"sync"
	"time"
)

// LogFunction For big messages, it can be more efficient to pass a function
// and only call it if the log level is actually enables rather than
// generating the log message and then checking if the level is enabled
type LogFunction func() []interface{}

type Logger struct {
	// The logs are `io.Copy`'d to this in a mutex. It's common to set this to a
	// file, or leave it default which is `os.Stderr`. You can also set this to
	// something more adventurous, such as logging to Kafka.
	Out io.Writer
	// Hooks for the logger instance. These allow firing events based on logging
	// levels and log entries. For example, to send errors to an error tracking
	// service, log to StatsD or dump the core on fatal errors.
	Hooks LevelHooks
	// All log entries pass through the formatter before logged to Out. The
	// included formatters are `TextFormatter` and `JSONFormatter` for which
	// TextFormatter is the default. In development (when a TTY is attached) it
	// logs with colors, but to a file it wouldn't. You can easily implement your
	// own that implements the `Formatter` interface, see the `README` or included
	// formatters for examples.
	Formatter Formatter

	// Flag for whether to log caller info (off by default)
	ReportCaller bool

	// The logging level the logger should log at. This is typically (and defaults
	// to) `logrus.Info`, which allows Info(), Warn(), Error() and Fatal() to be
	// logged.
	Level Level
	// Used to sync writing to the log. Locking is enabled by Default
	mu MutexWrap
	// Reusable empty entry
	entryPool sync.Pool
	// Function to exit the application, defaults to `os.Exit()`
	ExitFunc exitFunc
}

type exitFunc func(int)

type MutexWrap struct {
	lock     sync.Mutex
	disabled bool
}

func (mw *MutexWrap) Lock() {
	if !mw.disabled {
		mw.lock.Lock()
	}
}

func (mw *MutexWrap) Unlock() {
	if !mw.disabled {
		mw.lock.Unlock()
	}
}

func (mw *MutexWrap) Disable() {
	mw.disabled = true
}

// Creates a new logger. Configuration should be set by changing `Formatter`,
// `Out` and `Hooks` directly on the default logger instance. You can also just
// instantiate your own:
//
//    var log = &logrus.Logger{
//      Out: os.Stderr,
//      Formatter: new(logrus.TextFormatter),
//      Hooks: make(logrus.LevelHooks),
//      Level: logrus.DebugLevel,
//    }
//
// It's recommended to make this a global instance called `log`.
func New() *Logger {
	panic("not implemented...yet...")
	return nil
}

// WithField allocates a new entry and adds a field to it.
// Debug, Print, Info, Warn, Error, Fatal or Panic must be then applied to
// this new returned entry.
// If you want multiple fields, use `WithFields`.
func (logger *Logger) WithField(key string, value interface{}) *Entry {
	panic("not implemented...yet...")
	return nil
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func (logger *Logger) WithFields(fields Fields) *Entry {
	panic("not implemented...yet...")
	return nil
}

// Add an error as single field to the log entry.  All it does is call
// `WithError` for the given `error`.
func (logger *Logger) WithError(err error) *Entry {
	panic("not implemented...yet...")
	return nil
}

// Add a context to the log entry.
func (logger *Logger) WithContext(ctx context.Context) *Entry {
	panic("not implemented...yet...")
	return nil
}

// Overrides the time of the log entry.
func (logger *Logger) WithTime(t time.Time) *Entry {
	panic("not implemented...yet...")
	return nil
}

func (logger *Logger) Logf(level Level, format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Tracef(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Log(level Level, args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) LogFn(level Level, fn LogFunction) {
	panic("not implemented...yet...")
}

func (logger *Logger) Trace(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Debug(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Info(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Print(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Warn(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Warning(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Error(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Fatal(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Panic(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) TraceFn(fn LogFunction) {
	panic("not implemented...yet...")
}

func (logger *Logger) DebugFn(fn LogFunction) {
	panic("not implemented...yet...")
}

func (logger *Logger) InfoFn(fn LogFunction) {
	panic("not implemented...yet...")
}

func (logger *Logger) PrintFn(fn LogFunction) {
	panic("not implemented...yet...")
}

func (logger *Logger) WarnFn(fn LogFunction) {
	panic("not implemented...yet...")
}

func (logger *Logger) WarningFn(fn LogFunction) {
	panic("not implemented...yet...")
}

func (logger *Logger) ErrorFn(fn LogFunction) {
	panic("not implemented...yet...")
}

func (logger *Logger) FatalFn(fn LogFunction) {
	panic("not implemented...yet...")
}

func (logger *Logger) PanicFn(fn LogFunction) {
	panic("not implemented...yet...")
}

func (logger *Logger) Logln(level Level, args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Traceln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Debugln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Infoln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Println(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Warnln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Warningln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Errorln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Fatalln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Panicln(args ...interface{}) {
	panic("not implemented...yet...")
}

func (logger *Logger) Exit(code int) {
	panic("not implemented...yet...")
}

//When file is opened with appending mode, it's safe to
//write concurrently to a file (within 4k message on Linux).
//In these cases user can choose to disable the lock.
func (logger *Logger) SetNoLock() {
	panic("not implemented...yet...")
}

// SetLevel sets the logger level.
func (logger *Logger) SetLevel(level Level) {
	panic("not implemented...yet...")
}

// GetLevel returns the logger level.
func (logger *Logger) GetLevel() Level {
	panic("not implemented...yet...")
	return InfoLevel
}

// AddHook adds a hook to the logger hooks.
func (logger *Logger) AddHook(hook Hook) {
	panic("not implemented...yet...")
}

// IsLevelEnabled checks if the log level of the logger is greater than the level param
func (logger *Logger) IsLevelEnabled(level Level) bool {
	panic("not implemented...yet...")
	return false
}

// SetFormatter sets the logger formatter.
func (logger *Logger) SetFormatter(formatter Formatter) {
	panic("not implemented...yet...")
}

// SetOutput sets the logger output.
func (logger *Logger) SetOutput(output io.Writer) {
	panic("not implemented...yet...")
}

func (logger *Logger) SetReportCaller(reportCaller bool) {
	panic("not implemented...yet...")
}

// ReplaceHooks replaces the logger hooks and returns the old ones
func (logger *Logger) ReplaceHooks(hooks LevelHooks) LevelHooks {
	panic("not implemented...yet...")
	return nil
}
