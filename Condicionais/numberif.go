package main

import (
	"fmt"
)

func main() {

	var n1 int
	fmt.Println("Digite um número: ")
	fmt.Scan(&n1)

	if n1 == 100 {
		for i := 0; i < 100; i++ {
			fmt.Printf("%d\n", i)
		}
	} else if n1 == 50 {
		for i := 0; i < 50; i++ {
			fmt.Printf("%d\n", i)
		}
	} else {
		print(n1)
	}

}
