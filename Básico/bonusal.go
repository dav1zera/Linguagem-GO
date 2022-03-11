package main

import (
	"fmt"
)

func main() {

	var nome string
	var n1, n2, bonus float64

	fmt.Scanf("%s\n", &nome)
	fmt.Scanf("%f\n", &n1)
	fmt.Scanf("%f\n", &n2)
	bonus = (n1 + n2*0.15)

	fmt.Printf("TOTAL = R$ %.2f\n", bonus)

}
