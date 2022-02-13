package main

import "fmt"

func soma(n1 float64, n2 float64) float64 {
	return n1 + n2
}

func main() {

	var num1 float64
	var num2 float64
	var res float64
	fmt.Println("Digite um número: ")
	fmt.Scan(&num1)
	fmt.Println("Digite outro: ")
	fmt.Scan(&num2)

	res = soma(num1, num2)

	fmt.Println("X: ", res)

}
