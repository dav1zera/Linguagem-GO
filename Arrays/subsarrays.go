package main

import (
	"fmt"
)

func main() {

	var ini int64
	var ints = [10]int64{}
	for ini = 0; ini < 10; ini++ {
		fmt.Scanf("%i", &ints[ini])

	}

	for ini = 0; ini < 10; ini++ {
		if ints[ini] <= 0 {
			ints[ini] = 1
		}

	}

	for ini = 0; ini < 10; ini++ {
		fmt.Printf("[%d] = %d\n", ini, ints[ini])
	}

}
