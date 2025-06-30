package main

import (
	"fmt"
	"math"
)

/**
# Interface Segregation Principle (ISP)
- If one can offer a mechanism where clients are not dependent on interfaces they do not utilize, it should be encouraged.
- Somewhere these smaller specialized interfaces are preferred over having one larger general purpose interface.
- This means that clients use only those methods they need, to help them achieve their goals. Flux also helps minimize the effects of changes and increase flexibility.
**/

type square struct {
	length float64
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

type cube struct {
	square
}

func (c cube) volume() float64 {
	return math.Pow(c.length, 3)
}

type shape interface {
	area() float64
}

type object interface {
	shape
	volume() float64
}

func areaSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

func areaVolumeSum(shapes ...object) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return sum
}

func main() {
	s1 := square{length: 5}
	s2 := square{length: 6}
	c1 := cube{square: square{length: 3}}
	c2 := cube{square: square{length: 4}}
	fmt.Println(areaSum(s1, s2, c1, c2))
	fmt.Println(areaVolumeSum(c1, c2))
}
