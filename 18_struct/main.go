package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func main() {
	p1 := newPerson("Alice", 30)

	fmt.Println("Name:", p1.name)
	fmt.Println("Age:", p1.age)

	p1.age = 99
	fmt.Println("Updated Age:", p1.age)

	p2 := Person{name: "Bob"}

	fmt.Println("Name:", p2.name)
	fmt.Println("Age:", p2.age)
	fmt.Println("----------")

	fmt.Println("Age:", p1.age)

	addAgePtr(p1, 5)
	fmt.Println("Age after addAgePtr:", p1.age)

	fmt.Println("Age:", p2.age)

	addAgeValue(p2, 5)
	fmt.Println("Age after addAgeValue:", p2.age)
}

func addAgePtr(p *Person, years int) {
	p.age += years
}

func addAgeValue(p Person, years int) {
	p.age += years
}
