package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//qtd := 6
	//double(&qtd)
	//fmt.Printf("%d\n", qtd)

	qtd := 6
	fmt.Println(qtd)
	fmt.Println(&qtd)
	var end *int = &qtd

	fmt.Printf("O endereço é: %v com tamanho de %d bits\n", end, unsafe.Sizeof(end))
	//fmt.Println(end, unsafe.Sizeof(end))
}

func double(num *int) {
	*num *= 2
}
