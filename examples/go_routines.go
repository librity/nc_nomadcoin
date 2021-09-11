package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	go randomCounter("first")
	go randomCounter("second")

	wait()
}

func count(label string) {
	for i := range [10]int{} {
		fmt.Println(label, i)
		time.Sleep(1 * time.Second)
	}
}

func wait() {
	bufio.NewScanner(os.Stdin).Scan()
}
