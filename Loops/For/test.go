package main

import (
	"fmt"
)

func main() {

	var n1 int64
	fmt.Scanf("%i", &n1)

	for n1 = 0; n1 > 0; {
		n1++
		fmt.Println(n1)
		break
	}

	for n1 = 0; n1 < 0; {
		n1--
		fmt.Println(n1)
		break
	}

}
