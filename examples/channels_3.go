package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	routines   = 2
	countCap   = 2
	countTotal = routines * countCap
	interval   = time.Duration(1000) * time.Millisecond
)

func main() {
	chanDemo()

	wait()
}

func chanDemo() {
	c := make(chan int)

	launchRoutines(c)

	receiveAllAndClose(c)
}

// Send-only channel
func launchRoutines(c chan<- int) {
	// ERROR:
	// a := <-c

	for id := 0; id < routines; id++ {
		go count(id, c)
	}
}

func receiveAllAndClose(c chan int) {
	for i := 0; i < countTotal; i++ {
		receiveResultSafe(c)
	}

	// Cannot close a receive-only channel
	close(c)
	fmt.Println("MAIN: Closing channel.", c)
}

// Receive-only channel
func receiveResultSafe(c <-chan int) {
	// ERROR:
	// c <- 1

	fmt.Println("MAIN: Blocked, awaiting result...")

	result, ok := <-c
	if !ok {
		fmt.Println("MAIN: Unblocked, channel", c, "is closed")
		fmt.Println("---")
		return
	}

	fmt.Println("MAIN: Unblocked, result received.")
	fmt.Println("MAIN: result:", result)
	fmt.Println("---")
}

func count(id int, c chan<- int) {
	for i := 0; i < countCap; i++ {
		time.Sleep(interval)

		fmt.Println("ROUTINE", id, ": sending", i)
		c <- i
	}
}

func wait() {
	bufio.NewScanner(os.Stdin).Scan()
}
