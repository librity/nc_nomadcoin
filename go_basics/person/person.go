package person

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) SetDetailsBad(name string, age int) {
	p.name = name
	p.age = age

	fmt.Println("SetDetails' 'lior'", p)
}

func (p *Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age

	fmt.Println("SetDetails' 'lior'", p)
}

func (p Person) GetName() string {
	return p.name
}
