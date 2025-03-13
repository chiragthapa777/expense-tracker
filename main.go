package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) greet() {
	fmt.Println(p.Name)
}

type Human struct {
	Person
	Age string
}

func changeName(p *Human) {
	p.Name = "new Name"
}

func main() {
	h := Human{
		Age: "test",
		Person: Person{
			Name: "chirag",
		},
	}

	h.greet()

	changeName(&h)

	h.greet()
}
