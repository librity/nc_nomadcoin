package main

import "fmt"

func main() {

	demo1()
	fmt.Println("---")

	demo2()
	fmt.Println("---")

}

func demo1() {
	a := 2
	b := a
	a = 12

	fmt.Println(b)
	fmt.Println("a", &a, a)
	fmt.Println("b", &b, b)
}

func demo2() {
	a := 2
	b := &a
	a = 12
	a = 42

	fmt.Println("a", &a, a)
	fmt.Println("b", &b, b, *b)
	fmt.Println("-")

	a = -42
	fmt.Println("a", &a, a)
	fmt.Println("b", &b, b, *b)

}
