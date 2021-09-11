package main

// Both sending and receiveing are blocking operations

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	routines      = 2
	countCap      = 5
	countTotal    = routines * countCap
	countInterval = time.Duration(1000) * time.Millisecond
	receiveDelay  = time.Duration(10) * time.Second
)

func main() {
	chanDemo()

	wait()
}

func chanDemo() {
	// Buffered channel, only blocks after 5 unread sends.
	c := make(chan int, 3)

	launchRoutines(c)
	receiveAllAndClose(c)
}

func launchRoutines(c chan<- int) {
	for id := 0; id < routines; id++ {
		go count(id, c)
	}
}

func receiveAllAndClose(c chan int) {
	for i := 0; i < countTotal; i++ {
		receiveResultSafe(c)
	}

	close(c)
	fmt.Println("MAIN: Closing channel.", c)
}

func receiveResultSafe(c <-chan int) {
	fmt.Println("MAIN: Blocked, awaiting result...")

	time.Sleep(receiveDelay)
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
		// time.Sleep(countInterval)

		fmt.Println("ROUTINE", id, ": sending", i)
		c <- i
		fmt.Println("ROUTINE", id, ": sent", i)
	}
}

func wait() {
	bufio.NewScanner(os.Stdin).Scan()
}
