package pleat

import (
	"golang.org/x/sys/unix"
)

// isTerminal ...
func isTerminal(fd int) bool {
	_, err := unix.IoctlGetTermio(fd, unix.TCGETA)
	return err == nil
}
