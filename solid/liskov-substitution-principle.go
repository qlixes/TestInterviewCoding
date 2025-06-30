package main

import "fmt"

/**
# Liskov Substitution Principle (LSP)
- Subtypes must be able to substitute their base types and thereby not affect the integrity of the program in the process.
- Generated classes should degrade gracefully while adding functionality to the base class and not alter its behavior.
- It guarantees that you can freely use a derived class instead of the base class thereby improving reliability and robustness.
**/

type human struct {
	name string
}

func (h human) getName() string {
	return h.name
}

type teacher struct {
	human
	degree string
	salary float64
}

type student struct {
	human
	grades map[string]int
}

type person interface {
	getName() string
}

type printer struct {
}

func (printer) info(p person) {
	fmt.Println("Name: ", p.getName())
}

func main() {
	h := human{name: "Alex"}
	s := student{
		human: human{name: "Mike"},
		grades: map[string]int{
			"Math":    8,
			"English": 9,
		},
	}
	t := teacher{
		human:  human{name: "John"},
		degree: "CS",
		salary: 2000,
	}

	p := printer{}
	p.info(h)
	p.info(s)
	p.info(t)
}
