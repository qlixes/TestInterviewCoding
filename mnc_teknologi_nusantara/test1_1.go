package main

import (
	"fmt"
)

func test1(stacks []string) ([]int, bool) {
	var result []int
	var clean_map = make(map[string]string)
	var focus_item = ""

	for _, needle := range stacks {
		_, ok := clean_map[needle]

		if ok {
			focus_item = needle
			clean_map = make(map[string]string)
		} else {
			clean_map[needle] = needle
		}
	}

	for id_item, item := range stacks {
		if item == focus_item {
			result = append(result, id_item+1)
		}
	}

	if len(result) == 0 {
		return result, false
	}

	return result, true
}

func main() {
	list1 := []string{"abcd", "acbd", "aaab", "acbd"}

	res, status := test1(list1)

	if status {
		fmt.Printf("%v", res)
	} else {
		fmt.Println(status)
	}

	list2 := []string{
		"Satu",
		"Sate",
		"Tujuh",
		"Tusuk",
		"Tujuh",
		"Sate",
		"Bonus",
		"Tiga",
		"Puluh",
		"Tujuh",
		"Tusuk",
	}

	res2, status2 := test1(list2)

	if status2 {
		fmt.Printf("%v", res2)
	} else {
		fmt.Println(status2)
	}
}
