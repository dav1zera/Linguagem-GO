package main

import (
	"fmt"
	"math"
)

func main() {

	var nA, nB, nC, delta, resL1, resL2 float64
	fmt.Println("Digite 3 valores: ")
	fmt.Scanf("%f %f %f", &nA, &nB, &nC)

	if nA != 0 {
		delta = (nB * nB) - (4 * nA * nC)

		if delta < 0 {
			fmt.Printf("Não tem raíz real\n")
		} else if delta == 0 {
			resL1 = (-nB) / (2 * nA)
			fmt.Println("Possui apenas uma raíz real\n", resL1)
		} else {
			resL1 = (-nB - math.Sqrt(delta)/(2*nA))
			resL2 = (-nB + math.Sqrt(delta)/(2*nA))
			fmt.Printf("L1 = %f\n", resL1)
			fmt.Printf("L2 = %f\n", resL2)
		}

	} else {
		fmt.Printf("Não é equação do 2 grau\n")
	}

}
