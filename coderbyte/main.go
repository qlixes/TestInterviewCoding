package main

import (
	"regexp"
	"strings"
)

func LongestWord(sen string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	clean := re.ReplaceAllString(sen, "")

	stack := strings.Split(clean, " ")

	result := ""

	for _, j := range stack {
		if len(result) < len(j) {
			result = j
		}
	}

	return result
}

func FirstFactorial(num int) int {
	result := 1

	for i := 1; i <= num; i++ {
		result = result * i
	}

	return result
}

func main() {
}
