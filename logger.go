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
	panic("Logger.WithField not implemented...yet...")
	return nil
}

// WithFields ...
func (logger *Logger) WithFields(fields ...interface{}) *Entry {
	entry := NewEntry(logger)

	return entry.WithFields(fields...)
}

// WithError ...
func (logger *Logger) WithError(err error) *Entry {
	entry := NewEntry(logger)

	return entry.WithError(err)
}

// WithContext ...
func (logger *Logger) WithContext(ctx context.Context) *Entry {
	panic("Logger.WithContext not implemented...yet...")
	return nil
}

// WithTime ...
func (logger *Logger) WithTime(t time.Time) *Entry {
	panic("Logger.WithTime not implemented...yet...")
	return nil
}

func (logger *Logger) Logf(level Level, format string, args ...interface{}) {
	panic("Logger.Logf not implemented...yet...")
}

func (logger *Logger) Tracef(format string, args ...interface{}) {
	panic("Logger.Tracef not implemented...yet...")
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	panic("Logger.Debugf not implemented...yet...")
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	panic("Logger.Infof not implemented...yet...")
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	panic("Logger.Printf not implemented...yet...")
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	panic("Logger.Warnf not implemented...yet...")
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	panic("Logger.Warningf not implemented...yet...")
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	panic("Logger.Errorf not implemented...yet...")
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	panic("Logger.Fatalf not implemented...yet...")
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	panic("Logger.Panicf not implemented...yet...")
}

func (logger *Logger) Log(level Level, args ...interface{}) {
	if logger.IsLevelEnabled(level) {
		entry := logger.getEntry()
		entry.Info(args...)
		logger.putEntry(entry)
	}
}

func (logger *Logger) LogFn(level Level, fn LogFunction) {
	panic("Logger.LogFn not implemented...yet...")
}

func (logger *Logger) Trace(args ...interface{}) {
	logger.Log(TraceLevel, args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.Log(DebugLevel, args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.Log(InfoLevel, args...)
}

func (logger *Logger) Print(args ...interface{}) {
	logger.Info(args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.Log(WarnLevel, args...)
}

func (logger *Logger) Warning(args ...interface{}) {
	logger.Log(WarnLevel, args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.Log(ErrorLevel, args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.Log(FatalLevel, args...)
}

func (logger *Logger) Panic(args ...interface{}) {
	logger.Log(PanicLevel, args...)
}

func (logger *Logger) TraceFn(fn LogFunction) {
	panic("Logger.TraceFn not implemented...yet...")
}

func (logger *Logger) DebugFn(fn LogFunction) {
	panic("Logger.DebugFn not implemented...yet...")
}

func (logger *Logger) InfoFn(fn LogFunction) {
	panic("Logger.InfoFn not implemented...yet...")
}

func (logger *Logger) PrintFn(fn LogFunction) {
	panic("Logger.PrintFn not implemented...yet...")
}

func (logger *Logger) WarnFn(fn LogFunction) {
	panic("Logger.WarnFn not implemented...yet...")
}

func (logger *Logger) WarningFn(fn LogFunction) {
	panic("Logger.WarningFn not implemented...yet...")
}

func (logger *Logger) ErrorFn(fn LogFunction) {
	panic("Logger.ErrorFn not implemented...yet...")
}

func (logger *Logger) FatalFn(fn LogFunction) {
	panic("Logger.FatalFn not implemented...yet...")
}

func (logger *Logger) PanicFn(fn LogFunction) {
	panic("Logger.PanicFn not implemented...yet...")
}

func (logger *Logger) Logln(level Level, args ...interface{}) {
	panic("Logger.Logln not implemented...yet...")
}

func (logger *Logger) Traceln(args ...interface{}) {
	panic("Logger.Traceln not implemented...yet...")
}

func (logger *Logger) Debugln(args ...interface{}) {
	panic("Logger.Debugln not implemented...yet...")
}

func (logger *Logger) Infoln(args ...interface{}) {
	panic("Logger.Infoln not implemented...yet...")
}

func (logger *Logger) Println(args ...interface{}) {
	panic("Logger.Println not implemented...yet...")
}

func (logger *Logger) Warnln(args ...interface{}) {
	panic("Logger.Warnln not implemented...yet...")
}

func (logger *Logger) Warningln(args ...interface{}) {
	panic("Logger.Warningln not implemented...yet...")
}

func (logger *Logger) Errorln(args ...interface{}) {
	panic("Logger.Errorln not implemented...yet...")
}

func (logger *Logger) Fatalln(args ...interface{}) {
	panic("Logger.Fatalln not implemented...yet...")
}

func (logger *Logger) Panicln(args ...interface{}) {
	panic("Logger.Panicln not implemented...yet...")
}

func (logger *Logger) Exit(code int) {
	panic("Logger.Exit not implemented...yet...")
}

//When file is opened with appending mode, it's safe to
//write concurrently to a file (within 4k message on Linux).
//In these cases user can choose to disable the lock.
func (logger *Logger) SetNoLock() {
	panic("Logger.SetNoLock not implemented...yet...")
}

// SetLevel ...
func (logger *Logger) SetLevel(level Level) {
	logger.Level = level
}

// GetLevel ...
func (logger *Logger) GetLevel() Level {
	panic("Logger.GetLevel not implemented...yet...")
	return InfoLevel
}

// AddHook ...
func (logger *Logger) AddHook(hook Hook) {
	panic("Logger.AddHook not implemented...yet...")
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
	panic("Logger.ReplaceHooks not implemented...yet...")
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
