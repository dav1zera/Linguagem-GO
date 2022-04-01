package main

import (
	"fmt"
)

func main() {
	var n1 int64 = 0
	var sum int64 = 0
	var counter float64 = 0.0

	for {
		fmt.Scanf("%i", &n1)
		if n1 >= 0 {
			counter++
			sum += n1

		}
	}
}
