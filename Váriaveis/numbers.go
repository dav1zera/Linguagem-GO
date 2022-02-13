package main

import "fmt"

const ini = 50
const pi = 3.14

func main() {

	var num1 uint64
	var num2 uint32
	var num3 uint16
	var num4 float32
	var num5 float64
	var num6 complex128
	var num7 complex64

	num1 = 10
	num5 = 10.012122

	fmt.Printf("%v, %T\n", num7, num7)
	fmt.Printf("%v, %T\n", num6, num6)
	fmt.Printf("%v, %T\n", num5, num5)
	fmt.Printf("%v, %T\n", num4, num4)
	fmt.Printf("%v, %T\n", num3, num3)
	fmt.Printf("%v, %T\n", num2, num2)
	fmt.Printf("%v, %T\n", num1, num1)
	fmt.Printf("%p\n", &num1)

	fmt.Printf("%T", ini)

}
