package main

import "fmt"

func main() {
	x := 31883801889850185
	name := "Garmonbozia"

	printIntBinary(x)
	fmt.Println("---")

	printIntOctal(x)
	fmt.Println("---")

	printIntHex(x)
	fmt.Println("---")

	printIntHEX(x)
	fmt.Println("---")

	printIntUnicode(x)
	fmt.Println("---")

	printIntAndBinary(x)
	fmt.Println("---")

	printStrQuoted(name)
	fmt.Println("---")

	printStrHex(name)
	fmt.Println("---")

	printStrHEX(name)
	fmt.Println("---")

}

func printIntBinary(number int) {
	fmt.Printf("%b\n", number)
}

func printIntOctal(number int) {
	fmt.Printf("%o\n", number)
}

func printIntHex(number int) {
	fmt.Printf("%x\n", number)
}

func printIntHEX(number int) {
	fmt.Printf("%X\n", number)
}

func printIntUnicode(number int) {
	fmt.Printf("%U\n", number)
}

func printIntAndBinary(number int) {
	binary := fmt.Sprintf("%b", number)
	fmt.Println(number, binary)
}

func printStrQuoted(str string) {
	quoted := fmt.Sprintf("%q", str)
	fmt.Println(str, quoted)
}

func printStrHex(str string) {
	quoted := fmt.Sprintf("%x", str)
	fmt.Println(str, quoted)
}

func printStrHEX(str string) {
	quoted := fmt.Sprintf("%X", str)
	fmt.Println(str, quoted)
}
