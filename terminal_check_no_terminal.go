// +build js nacl plan9

package pleat

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
