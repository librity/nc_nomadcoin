package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	routines   = 3
	countCap   = 3
	countTotal = routines * countCap
	interval   = time.Duration(1000) * time.Millisecond
)

func main() {
	c := make(chan int)

	for id := 0; id < routines; id++ {
		go count(id, c)
	}

	for i := 0; i < countTotal; i++ {
		receiveResult(c)
	}

	fmt.Println("MAIN: Deadlock, will panic and exit!")
	receiveResult(c)

	wait()
}

func receiveResult(c chan int) {
	fmt.Println("MAIN: Blocked, awaiting result...")
	result := <-c
	fmt.Println("MAIN: Unblocked, result received.")
	fmt.Println("MAIN: result:", result)
	fmt.Println("---")
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
