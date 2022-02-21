package main

import (
	"fmt"
)

type Multi struct {
	n1 float64
	n2 float64
}

func main() {

	var num1, num2, res float64
	fmt.Println(" ")
	fmt.Scan(&num1)
	fmt.Println(" ")
	fmt.Scan(&num2)

	mul := Multi{
		n1: num1,
		n2: num2,
	}

	res = mul.n1 * mul.n2

	fmt.Println(res)
}
