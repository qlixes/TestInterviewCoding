package coderbyte

import (
	"regexp"
	"strings"
)

func LongestWord(sen string) string {

	re := regexp.MustCompile(`[^a-zA-Z0-9 ]`)
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

}
