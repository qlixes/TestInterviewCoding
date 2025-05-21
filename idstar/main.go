package main

import (
	"fmt"
	"strconv"
	"strings"
)

func RotateNum(stack []int) string {

	size := len(stack)
	result := make([]string, size)

	mode := stack[0]
	id := 0

	if len(stack) < mode {
		return strings.Join(result, "")
	}

	for i, j := range stack {

		if i < mode {
			id = size - mode + i
		}

		if i >= mode {
			id = i - mode
		}

		result[id] = strconv.Itoa(j)
	}

	return strings.Join(result, "")
}

func StyleWord(word string, size int) [][]byte {

	cols := make([]byte, len(word))
	result := make([]cols, size)

	x, y := 0, 0

	for i, j := range word {
		x += 1
		y += 1
	}

	fmt.Println(result)
}

func main() {
	fmt.Println(RotateNum([]int{3,4,6,7,1,5,6,7,8,9}))
	StyleWord("wordpress", 5)
}