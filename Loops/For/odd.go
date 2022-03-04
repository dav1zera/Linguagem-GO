package main

import (
	"fmt"
)

func main() {

	var num int64
	fmt.Scanf("%d", num)

	for count := 1; count <= 100; count++ {
		if count%2 != 0 {
			fmt.Printf("%d\n", count)
		}
	}

}
