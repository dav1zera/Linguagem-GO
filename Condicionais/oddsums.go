package main

import (
	"fmt"
)

func main() {
	var tmp = 0

	var n1, n2, min, max int64

	fmt.Scanf("%d", &n1)
	fmt.Scanf("%d", &n2)

	if n1 < n2 {
		min = n1
		max = n2
	} else {
		max = n2
		min = n1
	}

	for count := (min + 1); count < max; count++ {
		if count%2 == 1 || count%2 == -1 {
			tmp += 1
		}
	}

	fmt.Printf("%d", tmp)

}
