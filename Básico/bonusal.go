package main

import (
	"fmt"
)

func main() {

	var nome string
	var n1, n2, bonus float64

	fmt.Println("Digite o nome do funcionário: ")
	fmt.Scanf("%s", &nome)

	fmt.Scanf("%f %f", &n1, &n2)
	bonus = (n1 + n2*0.15)

	fmt.Printf("Bonus = R$ %.2f\n", bonus)

}
