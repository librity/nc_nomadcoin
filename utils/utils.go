// Package utils implements misc functions to be used by all other packages.
package utils

import (
	"bufio"
	"os"
)

// Wait pauses the current go routine until STDIN receives a new line.
func Wait() {
	bufio.NewScanner(os.Stdin).Scan()
}
