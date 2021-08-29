package main

import (
	"fmt"

	"github.com/librity/nc_nomadcoin/go_basics/person"
)

func main() {

	demo1()
	fmt.Println("---")

	demo2()
	fmt.Println("---")

	demo3()
	fmt.Println("---")

}

func demo1() {
	lior := person.Person{}

	fmt.Println("Main's 'lior'", lior)
}

func demo2() {
	lior := person.Person{}
	lior.SetDetailsBad("Lior", 18)

	fmt.Println("Main's 'lior'", lior)
}

func demo3() {
	lior := person.Person{}
	lior.SetDetails("Lior", 18)

	fmt.Println("Main's 'lior'", lior)
	fmt.Println("name:", lior.GetName())
}
