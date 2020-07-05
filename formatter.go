package pleat

import (
	"time"
)

// Default key names ...
const (
	defaultTimestampFormat = time.RFC3339
	FieldKeyMsg            = "msg"
	FieldKeyLevel          = "level"
	FieldKeyTime           = "time"
	FieldKeyLogrusError    = "logrus_error"
	FieldKeyFunc           = "func"
	FieldKeyFile           = "file"
)

// Formatter ...
type Formatter interface {
	Format(*Entry) ([]byte, error)
}
