package main

import (
	"fmt"
)

func getYearsLeft(p, r, f int) int {
	years := 0
	for p <= f {
		years++
		p = p * r
	}
	return years
}

func main() {
	sample1 := getYearsLeft(1, 3, 9)
	fmt.Println(sample1)
	sample2 := getYearsLeft(2, 2, 16)
	fmt.Println(sample2)
	sample3 := getYearsLeft(5, 2, 11)
	fmt.Println(sample3)
}
