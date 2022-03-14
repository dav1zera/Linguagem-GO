package main

import (
	"fmt"
	"math"
)

func max(nums ...int) int {
	max := int(math.Inf(-1))
	for _, n := range nums {
		if n > max {

			max = n
		}
	}
	return max
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Printf("Maior: %d", max(numbers...))
}
