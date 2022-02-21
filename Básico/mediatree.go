package main

import (
	"fmt"
)

func main() {

	var n1, n2, n3, media float64
	fmt.Scanf("%f %f %f", &n1, &n2, &n3)

	media = (n1*2 + n2*3 + n3*5) / (2 + 3 + 5)

	fmt.Printf("A media é: %.2f", media)

}
