package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ArrayChallenge02(stack []int) string {

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

func SearchingChallange(str string) string {
	// remove whitespace
	clean := strings.ReplaceAll(str, " ", "")

	help := make(map[rune]int)

	var sort []rune
	var result []string

	for _, j := range clean {
		help[j] += 1

		if help[j] == 1 {
			sort = append(sort, j)
		}
	}

	for _, j := range sort {
		if help[j] == 1 {
			result = append(result, string(j))
		}
	}

	return result[0]
}

func StringChallenge(word string, size int) string {

	return ""
}

func main() {
	fmt.Println(ArrayChallenge02([]int{3, 4, 6, 7, 1, 5, 6, 7, 8, 9}))
	fmt.Println(SearchingChallange("hello world hi hey"))
}
