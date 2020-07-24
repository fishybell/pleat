package pleat

import (
	"runtime"
)

type FieldKey string

// FieldMap ...
type FieldMap map[FieldKey]string

func (f FieldMap) resolve(key FieldKey) string {
	if k, ok := f[key]; ok {
		return k
	}

	return string(key)
}

// JSONFormatter ...
type JSONFormatter struct {
	TimestampFormat   string
	DisableTimestamp  bool
	DisableHTMLEscape bool
	DataKey           string
	FieldMap          FieldMap
	CallerPrettyfier  func(*runtime.Frame) (function string, file string)
	PrettyPrint       bool
}

// Format ...
func (f *JSONFormatter) Format(entry *Entry) ([]byte, error) {
	panic("JSONFormatter.Format not implemented...yet...")
	return nil, nil
}
