package main

import (
	"fmt"
)

func main() {
	var N int64
	var I int64

	var menor int64 = 0
	var posMenor int64 = 0
	fmt.Scanf("%d\n", &N)
	X := make([]int64, N)

	for I = 0; I < N; I++ {
		fmt.Scanf("%d", &X[I])

	}
	menor = X[0]
	for I = 0; I < N; I++ {

		if X[I] < menor {
			menor = X[I]
			posMenor = I
		}

	}
	fmt.Printf("Menor valor: %d\n", menor)
	fmt.Printf("Posicao: %d\n", posMenor)

}
