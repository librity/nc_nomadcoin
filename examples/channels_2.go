package main

// - Reading from a channel without active go routines will create a panic
// - Reading from a closed channel will return null
// - Closing a closed channel will create a panic
// - Sending to a closed channel will create a panic

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
	// go chanDemo()
	// go chanDemo()
	// chanDemo()
	chanDemo()

	wait()
}

func chanDemo() {
	c := make(chan int)

	launchRoutines(c)
	// receiveAllAndClose(c)
	// receiveResultSafe(c)
	deadlockAndDie(c)
	deadlockAndDie(c)
	// sendAndDie(c)
	// closeAndDie(c)
}

func launchRoutines(c chan int) {
	for id := 0; id < routines; id++ {
		go count(id, c)
	}
}

func receiveAllAndClose(c chan int) {
	for i := 0; i < countTotal; i++ {
		receiveResultSafe(c)
	}

	fmt.Println("MAIN: Closing channel.", c)
	close(c)
}

func receiveResultSafe(c chan int) {
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

func deadlockAndDie(c chan int) {
	fmt.Println("MAIN: Deadlock, will panic and exit unless channel", c, "has been closed")
	receiveResult(c)
}

func receiveResult(c chan int) {
	fmt.Println("MAIN: Blocked, awaiting result...")
	result := <-c
	fmt.Println("MAIN: Unblocked, result received.")
	fmt.Println("MAIN: result:", result)
	fmt.Println("---")
}

func closeAndDie(c chan int) {
	fmt.Println("MAIN: Will panic and exit if channel", c, "has been closed already")
	close(c)
}

func sendAndDie(c chan int) {
	fmt.Println("MAIN: Will panic and exit if channel", c, "has been closed already")
	c <- 1
}

func count(id int, c chan int) {
	for i := 0; i < countCap; i++ {
		time.Sleep(interval)

		fmt.Println("ROUTINE", id, ": sending", i)
		c <- i
	}
}

func wait() {
	bufio.NewScanner(os.Stdin).Scan()
}
