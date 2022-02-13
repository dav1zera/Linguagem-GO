package main

import (
	"fmt"
)

// := tipo inferido

func main() {

	var nome string
	var num int
	var num1 uint64
	var teste bool

	teste = false
	nome = "Davi"
	num1 = 10.0000
	num = 10

	fmt.Println("", teste)
	fmt.Println("", num)
	fmt.Println("", num1)
	fmt.Println("", nome)

	var i float64 = 42.5
	fmt.Printf("%v, %T\n", i, i)
	// => %v value
	// => %T type

	var isStatus bool
	isStatus = true

	fmt.Printf("%v, %T\n", isStatus, isStatus)

	a := 10
	b := 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a * b)
	fmt.Println(a % b)

	//Complex64

	var compl complex128
	fmt.Printf("%v, %T\n", compl, compl)
	fmt.Printf("%v, %T\n", compl, compl)
	fmt.Printf("%p", &compl) // => ponteiros

	var uns uint64
	fmt.Printf("%v, %T\n", uns, uns)

}
