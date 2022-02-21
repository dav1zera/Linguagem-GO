package main

import (
	"fmt"
)

func main() {

	var A int
	var B int

	fmt.Println("Digite o primeiro valor: ")
	fmt.Scanf("%d", &A)
	fmt.Println("Digite o segundo valor: ")
	fmt.Scanf("%d", &B)

	fmt.Printf("SOMA = %d\n", A+B)

}
