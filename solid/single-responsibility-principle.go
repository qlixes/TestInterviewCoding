package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

/**
# Single Responsibility Principle (SRP)
- Every class should have only one reason to change, which means that a class should have only one responsibility or its functionality should be limited to just one task.
- This makes the class easier to comprehend, evaluate, and build as compared to other programming structures that may be more complex to handle.
- It minimizes the level of code which enhances clarity, and it enhances the maintainability of the source code.
- This reduces the exposure of the system to bugs that cause changes because only fewer classes are affected.
**/

type circle struct {
	radius float64
}

func (c circle) name() string {
	return "circle"
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	length float64
}

func (s square) name() string {
	return "square"
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
	name() string
}

type outputter struct {
}

func (o outputter) Text(s shape) string {
	return fmt.Sprintf("area  of the %s: %f", s.name(), s.area())
}

func (o outputter) JSON(s shape) string {
	res := struct {
		Name string  `json:"shape"`
		Area float64 `json:"area"`
	}{
		Name: s.name(),
		Area: s.area(),
	}

	bs, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

func main() {
	c := circle{radius: 5}
	s := square{length: 7}
	out := outputter{}
	fmt.Println(out.Text(c))
	fmt.Println(out.JSON(c))
	fmt.Println(out.Text(s))
	fmt.Println(out.JSON(s))
}
