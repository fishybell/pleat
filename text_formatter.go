package pleat

import (
	"runtime"
)

// TextFormatter ...
type TextFormatter struct {
	ForceColors               bool
	DisableColors             bool
	ForceQuote                bool
	DisableQuote              bool
	EnvironmentOverrideColors bool
	DisableTimestamp          bool
	FullTimestamp             bool
	TimestampFormat           string
	DisableSorting            bool
	SortingFunc               func([]string)
	DisableLevelTruncation    bool
	PadLevelText              bool
	QuoteEmptyFields          bool
	FieldMap                  FieldMap
	CallerPrettyfier          func(*runtime.Frame) (function string, file string)
}

// Format ...
func (f *TextFormatter) Format(entry *Entry) ([]byte, error) {
	panic("TextFormatter.Format not implemented...yet...")
	return nil, nil
}
