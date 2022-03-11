package main

import (
	"fmt"
)

func main() {
	var nota1, nota2, media float64
	fmt.Scanf("%f\n", &nota1)
	fmt.Scanf("%f\n", &nota2)

	media = (nota1*3.5 + nota2*7.5) / (3.5 + 7.5)

	fmt.Printf("MEDIA = %.5f\n", media)
}
