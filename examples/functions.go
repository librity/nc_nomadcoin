package main

import "fmt"

func main() {
	result := plus(2, 3)
	fmt.Println(result)
	fmt.Println("---")

	result, name := plusSigned(4, 16, "Garmonbozia")
	fmt.Println(result, name)
	fmt.Println("---")

	result = plusMany(4, 15, 2, 5, 7, 1, 64, 12, 23, 58, 19)
	fmt.Println(result)
	fmt.Println("---")

	name = "Lior! ! ! ! ! ! Is my name! ! !  !! ! !!"
	analyzeString(name)
	fmt.Println("---")

	printString(name)
	fmt.Println("---")

	printBinary(name)
	fmt.Println("---")

	printOctal(name)
	fmt.Println("---")

	printHex(name)
	fmt.Println("---")

	printDecimal(name)
	fmt.Println("---")
}

func plus(a int, b int) int {
	return a + b
}

func plusSigned(a int, b int, name string) (int, string) {
	return a + b, name
}

func plusMany(numbers ...int) int {
	// total := 0
	var total int
	for _, value := range numbers {
		total += value
		// fmt.Println(index)
	}

	return total
}

func analyzeString(str string) {
	for index, byte := range str {
		fmt.Println(index, byte)
	}
}

func printString(str string) {
	for _, byte := range str {
		fmt.Print(string(byte))
	}
	fmt.Print("\n")
}

func printOctal(str string) {
	for _, byte := range str {
		fmt.Printf("%o ", byte)
	}
	fmt.Print("\n")
}

func printBinary(str string) {
	for _, byte := range str {
		fmt.Printf("%b ", byte)
	}
	fmt.Print("\n")
}

func printHex(str string) {
	for _, byte := range str {
		fmt.Printf("%x ", byte)
	}
	fmt.Print("\n")
}

func printDecimal(str string) {
	for _, byte := range str {
		fmt.Printf("%d ", byte)
	}
	fmt.Print("\n")
}
