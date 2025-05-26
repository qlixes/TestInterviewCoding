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

	return ""
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

func ArrayChallenge01(str []string) string {

	if len(str) > 2 {
		return ""
	}

	list := strings.ReplaceAll(str[0], ",", "")
	stack := strings.ReplaceAll(str[1], ",", "")

	duplicate := make(map[string]int)
	for i, j := range list {
		if i < len(stack) && string(stack[i]) == string(j) {
			duplicate[string(j)] += 1
		}
	}

	result := ""
	for i, j := range duplicate {
		if j == 1 {
			result += fmt.Sprint(i)
		}
		result += fmt.Sprint(",")
	}

	return result
}

func MergeSortedArrayChallenge(nums1 []int, m int, nums2 []int, n int) []int {
	var result []int

	// get list items based on m item
	slice1 := nums1[:m]
	slice2 := nums2[:n]

	// merge slice
	slice := append(slice1, slice2...)

	var helper []int

	temp := 0
	for i, j := range slice {
	}

	fmt.Println(helper)

	return result
}

func main() {
	fmt.Println(ArrayChallenge02([]int{4, 3, 4, 3, 1, 2}))
	fmt.Println(SearchingChallange("hello world hi hey"))
	fmt.Println(ArrayChallenge01([]string{"1,3,9,10,17,18", "1,4,9,10"}))
	fmt.Println(MergeSortedArrayChallenge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
}
