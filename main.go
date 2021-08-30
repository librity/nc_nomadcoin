package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data         string
	hash         string
	previousHash string
}

func main() {
	welcome()
	// hashDemo()
	bootstrap()
}

func welcome() {
	fmt.Println("Welcome to Nomad Coin!")
	fmt.Println("---")
}

func bootstrap() {
	genesisBlock := block{
		data:         "노마드 코인 - Fírst blóck in the chain",
		hash:         "",
		previousHash: ""}

	// inspectString(genesisBlock.data)
	genesisBlock.hash = hash(genesisBlock)
	fmt.Println(genesisBlock)
}

func inspectString(str string) {
	for _, character := range str {
		fmt.Printf("%b ", character)
	}
	fmt.Println()
	fmt.Println("---")
}

func hashDemo() {
	sum := sha256.Sum256([]byte("hello world\n"))
	fmt.Printf("%x\n", sum)
	fmt.Printf("%X\n", sum)
	fmt.Println("---")
}

func hash(b block) string {
	preHash := b.data + b.previousHash
	preHashBytes := []byte(preHash)
	rawHash := sha256.Sum256(preHashBytes)
	hexHash := fmt.Sprintf("%x", rawHash)

	fmt.Println(rawHash)
	fmt.Println("---")
	fmt.Println(hexHash)
	fmt.Println("---")

	return hexHash
}
