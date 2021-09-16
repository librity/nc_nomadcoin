package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/librity/nc_nomadcoin/utils"
)

const (
	countCap    = 20
	durationCap = 1000
)

func main() {
	labels := [...]string{"first", "second", "third"}
	channel := make(chan string)

	for _, label := range labels {
		go randomCounter(label, channel)
	}

	for _ = range labels {
		result := <-channel
		fmt.Println(result)
	}

	wait()
}

func randomCounter(label string, channel chan string) {
	utils.SeedRNG()
	randomCount := rand.Intn(countCap)
	utils.SeedRNG()
	randomDuration := rand.Intn(durationCap)
	interval := time.Duration(randomDuration) * time.Millisecond

	for i := 0; i < randomCount; i++ {
		fmt.Println(label, i)
		time.Sleep(interval)
	}

	channel <- label + " finished counting!"
}

func wait() {
	bufio.NewScanner(os.Stdin).Scan()
}
