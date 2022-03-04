package main

import (
	"fmt"
)

func main() {

	var n1, n2 int64
	fmt.Scanf("%d, %d", &n1, &n2)
	for n1 != n2 {

		if n1 > n2 {
			fmt.Printf("Crescente\n")
			break
		} else if n1 < n2 {
			fmt.Printf("Decrescente\n")
			break
		}

	}

}
