/*
Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: s = "cbaebabacd", p = "abc"
Output:
[0,6]

Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".

Example 2:
Input: s = "cbaebabacd", p = "aba"
Output:
[5]

Explanation:
The substring with start index = 5 is "aba", which is an anagram of "aba".

Example 3:
Input: s = "cbaebabacd", p = "babe"
Output:
[1,3]

Explanation:
The substring with start index = 1 is "baeb", which is an anagram of "babe".
The substring with start index = 3 is "ebab", which is an anagram of "babe".

Constraints:
len(s) >= len(p)
*/
package main

import (
	"fmt"
)

var s string
var p string

func anagram(dictionary string, needle string) []int {
	var result []int
	var stack []any

	// for pos, char := range s {
	// 	for id, nchar := range stack {
	// 		fmt.Printf("%d.%d [%v %v] ", pos, id, string(char), nchar)
	// 		if string(char) == nchar {
	// 			stack = removeKey(stack, id)

	// 			fmt.Print(" ✅ \n")
	// 			break
	// 		}
	// 		fmt.Print(" ❌ \n")
	// 	}
	// 	if len(stack) == 0 {
	// 		stack = strToArray(p)
	// 	}
	// }

	for pos, char := range dictionary {
		for _, nchar := range needle {
			fmt.Printf("%d [%v]", pos, string(char))
			// check increment
			if char == nchar {

				stack = append(stack, char)

				if pos == 0 {
					result = append(result, pos)
				}

				fmt.Print(" ✅ \n")
				break
			}
			fmt.Print(" ❌ \n")
		}
		// check increment
		if pos > 0 {
			if stack[pos-1] != nil {

			}
		}
	}

	return result
}

// func strToArray(s string) []any {
// 	var stack []any

// 	for _, val := range s {
// 		stack = append(stack, string(val))
// 	}

// 	return stack
// }

// // TODO : increase loop time using conccurency
// func removeKey(s []any, id int) []any {
// 	var result []any

// 	for i, val := range s {
// 		if i == id {
// 			continue
// 		}
// 		result = append(result, val)
// 	}

// 	return result
// }

func main() {
	s = "cbaebabacd"
	p = "abc"

	fmt.Println(anagram(s, p))

	s = "cbaebabacd"
	p = "aba"

	fmt.Println(anagram(s, p))

	s = "cbaebabacd"
	p = "babe"

	fmt.Println(anagram(s, p))
}
