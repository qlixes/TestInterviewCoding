package main

import (
	"fmt"
	"math"
)

/**
# Open-Closed Principle (OCP)
- Organizations and systems should be designed to be easily extensible but on the other hand, the existing interfaces should not be alterable or modifiable easily.
- One circuit that is unique with modules is that the action of a module can be further mannered without a direct signal to its source program.
- This is mainly done through the process of abstraction and polymorphism. It promotes flexibility and eliminates the possibility of having bugs creeping in as developers add new features.
**/

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) area() float64 {
	return t.base * t.height / 2
}

type shape interface {
	area() float64
}

type calculator struct {
}

func (a calculator) areaSum(shapes ...shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.area()
	}
	return sum
}

func main() {
	c := circle{radius: 5}
	s := square{length: 7}
	t := triangle{height: 3, base: 7}
	calc := calculator{}
	fmt.Println("area sum:", calc.areaSum(c, s, t))
}
