// +build appengine

package pleat

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
