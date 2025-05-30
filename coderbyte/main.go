package main

import (
	"fmt"
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

func main() {
	fmt.Println(LongestWord("123456789 98765432"))
}
