package main

import "fmt"

// Illegal:
// NAME := "Lior"
var NAME string = "LIOR"

const NAME2 string = "Garmonbozia"

func main() {
	string1()
	fmt.Println("---")

	string2()
	fmt.Println("---")

	string3()
	fmt.Println("---")
}

func string1() {
	var name string = "Lior"
	fmt.Println(name)

	// Illegal:
	// name = 12
	name = "12"
	fmt.Println(name)
}

func string2() {
	name := "Lior"
	fmt.Println(name)

	// Illegal:
	// name = 12
	name = "aa2asdaw"
	fmt.Println(name)
}

func string3() {
	fmt.Println(NAME)
	fmt.Println(NAME2)
}
