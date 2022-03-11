package main

import (
	"fmt"
)

func main() {

	var n1, n2, n3, media float64
	fmt.Scanf("%f\n", &n1)
	fmt.Scanf("%f\n", &n2)
	fmt.Scanf("%f\n", &n3)

	media = (n1*2 + n2*3 + n3*5) / (2 + 3 + 5)

	fmt.Printf("MEDIA = %.1f", media)

}
