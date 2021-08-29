package main

import (
	"fmt"
	"time"
)

type person struct {
	name      string
	birthYear int
}

func (p person) getAge() int {
	currentYear := int(time.Now().Year())

	return currentYear - p.birthYear
}

func (p person) getKoreanAge() int {
	currentYear := int(time.Now().Year())

	return currentYear - p.birthYear + 1
}

func (p person) sayHello() {
	fmt.Printf("Hello, my name is %s, I'm %d years old. \n",
		p.name, p.getAge())
}

func (p person) sayHelloInKorean() {
	fmt.Printf("Hello. My name is %s. My korean age is %d years old. \n",
		p.name, p.getKoreanAge())
}

func main() {

	demo1()
	fmt.Println("---")

}

func demo1() {
	Bolívar := person{name: "Simón Bolívar", birthYear: 1810}

	Bolívar.sayHello()
	Bolívar.sayHelloInKorean()
}
