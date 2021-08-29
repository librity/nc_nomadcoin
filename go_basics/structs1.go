package main

import "fmt"

type person struct {
	name string
	age  int
}

// Convention: use the first letter of the struct's name
func (p person) sayHi() {
	fmt.Printf("Hello, my name is %s and I am %d\n", p.name, p.age)
}

func (per person) inspect() {
	fmt.Printf("person{name: \"%s\", age: %d}\n", per.name, per.age)
}

func main() {

	demo1()
	fmt.Println("---")

}

func demo1() {
	nico := person{"Nico", 22}
	lior := person{name: "Lior", age: 18}

	nico.sayHi()
	lior.inspect()
}
