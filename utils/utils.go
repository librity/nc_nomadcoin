package utils

import (
	"bufio"
	"os"
)

func Wait() {
	bufio.NewScanner(os.Stdin).Scan()
}
