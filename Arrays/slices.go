package main

import "fmt"

func main() {
	//var nomes = []string{"Davi", "Alves", "Dos", "Santos"}

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("tamanho = %d\n", len(myslice))
	fmt.Printf("capacidade = %d\n", cap(myslice))

	//make()
	slints := make([]int, 5, 10)
	fmt.Printf("slints = %v\n", slints)
	fmt.Printf("tamanho = %d\n", len(slints))
	fmt.Printf("capacidade = %d\n", cap(slints))

}
