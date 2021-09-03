package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/librity/nc_nomadcoin/utils"
)

const (
	letters    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	verbose    = true
	difficulty = 2
)

var (
	target    = strings.Repeat("0", difficulty)
	blockData = randomString(20)
	nOnce     = 0
	attempts  = 1
	solution  string
)

func main() {
	greeting()
	mine()
	results()
}

func greeting() {
	fmt.Println("=== Proof of Work simulator ===")
	fmt.Println("Difficulty:", difficulty)
	fmt.Println("Block Data:", blockData)
}

func mine() {
	fmt.Println("Mining...")
	for {
		solution = utils.HexHash(blockData + fmt.Sprint(nOnce))

		if foundSolution() {
			break
		}
		handleLogging()

		nOnce++
		attempts++
	}
}

func foundSolution() bool {
	return strings.HasPrefix(solution, target)
}

func handleLogging() {
	if verbose {
		fmt.Println(attempts, solution)
	}
}

func results() {
	fmt.Println("Successfully mined block!")
	fmt.Println("NOnce:", nOnce)
	fmt.Println("Hash:", solution)
	fmt.Println("Attempts:", attempts)
}

func randomString(length int) string {
	seedRandom()

	randStr := make([]byte, length)
	for i := range randStr {
		random := rand.Intn(len(letters))
		randStr[i] = letters[random]
	}

	return string(randStr)
}

func seedRandom() {
	rand.Seed(time.Now().UnixNano())
}
