package main

import "fmt"

/*
* The Utopian Tree goes through 2 cycles of growth every year. Each spring, it doubles in height. Each summer, its height increases by 1 meter.
* A Utopian Tree sapling with a height of 1 meter is planted at the onset of spring. How tall will the tree be after  growth cycles?
 */
func utopianTree(n int32) int32 {
	spring := int32(2)
	summer := int32(1)
	result := int32(0)

	if n == 0 {
		return int32(1)
	}

	for i := int32(0); i <= n; i++ {
		if i%2 == 0 {
			result = result + summer
		}

		if i%2 > 0 {
			result = result * spring
		}
	}

	return result
}

func main() {
	fmt.Println(utopianTree(5))
}
