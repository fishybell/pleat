package pleat

import (
	"context"
	"io"
	"os"
	"sync"
	"time"
)

// LogFunction ...
type LogFunction func() []interface{}

type Logger struct {
	Out          io.Writer
	ErrorOut     io.Writer
	Hooks        LevelHooks
	Formatter    Formatter
	ReportCaller bool
	Level        Level
	ExitFunc     exitFunc

	mu        MutexWrap
	entryPool sync.Pool
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

// New ...
func New() *Logger {
	logger := &Logger{
		Out:      os.Stdout,
		ErrorOut: os.Stderr,
	}

	logger.entryPool = sync.Pool{
		New: func() interface{} {
			return NewEntry(logger)
		},
	}

	return logger
}

// WithField ...
func (logger *Logger) WithField(key string, value interface{}) *Entry {
	panic("not implemented...yet...")
	return nil
}

// WithFields ...
func (logger *Logger) WithFields(fields ...interface{}) *Entry {
	entry := NewEntry(logger)

	return entry.WithFields(fields...)
}

// WithError ...
func (logger *Logger) WithError(err error) *Entry {
	panic("not implemented...yet...")
	return nil
}

// WithContext ...
func (logger *Logger) WithContext(ctx context.Context) *Entry {
	panic("not implemented...yet...")
	return nil
}

// WithTime ...
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
	if logger.IsLevelEnabled(level) {
		entry := logger.getEntry()
		entry.Info(args...)
		logger.putEntry(entry)
	}
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
	logger.Log(InfoLevel, args...)
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

// SetLevel ...
func (logger *Logger) SetLevel(level Level) {
	logger.Level = level
}

// GetLevel ...
func (logger *Logger) GetLevel() Level {
	panic("not implemented...yet...")
	return InfoLevel
}

// AddHook ...
func (logger *Logger) AddHook(hook Hook) {
	panic("not implemented...yet...")
}

// IsLevelEnabled ...
func (logger *Logger) IsLevelEnabled(level Level) bool {
	return level <= logger.Level
}

// SetFormatter ...
func (logger *Logger) SetFormatter(formatter Formatter) {
	logger.Formatter = formatter
}

// SetOutput ...
func (logger *Logger) SetOutput(output io.Writer) {
	logger.Out = output
}

// SetErrorOutput ...
func (logger *Logger) SetErrorOutput(output io.Writer) {
	logger.ErrorOut = output
}

// SetReportCaller ...
func (logger *Logger) SetReportCaller(reportCaller bool) {
	logger.ReportCaller = reportCaller
}

// ReplaceHooks ...
func (logger *Logger) ReplaceHooks(hooks LevelHooks) LevelHooks {
	panic("not implemented...yet...")
	return nil
}

func (logger *Logger) getEntry() *Entry {
	entry, ok := logger.entryPool.Get().(*Entry)
	if ok {
		return entry
	}

	return NewEntry(logger)
}

func (logger *Logger) putEntry(entry *Entry) {
	entry.Data = Fields{}
	logger.entryPool.Put(entry)
}
