package main

import "fmt"

func main() {

	var number int64
	var qtdh float64
	var functrab, sal float64

	fmt.Scanf("%d %f %f", &number, &qtdh, &functrab)
	fmt.Printf("Número do funcionário: %d\n", number)

	sal = (qtdh * functrab)

	fmt.Printf("Salário: $US %.2f\n", sal)

}
