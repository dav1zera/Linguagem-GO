package main

import (
	"fmt"
)

const n = 3.14159

func areaCirc(value float64) float64 {
	return n * value * value
}

func main() {

	var area float64
	var number1 float64

	fmt.Scanf("%f\n", &number1)

	area = areaCirc(number1)

	fmt.Printf("A=%.4f\n", area)

}
