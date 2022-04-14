package main

import (
	"fmt"
)

func main() {

	notas := map[string]float64{
		"W": 1.0,
		"H": 1 / 2.0,
		"Q": 1 / 4.0,
		"E": 1 / 8.0,
		"S": 1 / 16.0,
		"T": 1 / 32.0,
		"X": 1 / 64.0,
	}

	var composiçãoMuscial string
	var contarComposto int64
	var soma float64

	var guardarNotas float64 = 0.0

	var i int64 = 0
	for i == 0 {
		fmt.Scanf("%s", &composiçãoMuscial)

		if composiçãoMuscial == "*" {
			break
		}

		soma = 0
		contarComposto = 0

		for j := 0; j < len(composiçãoMuscial); j++ {

			if string(composiçãoMuscial[j]) == "/" {
				if soma == 1 {
					contarComposto += 1
				}
				soma = 0
			} else {
				guardarNotas = notas[string(composiçãoMuscial[j])]

				if guardarNotas == 0 {
					soma = 2
				} else {
					soma += guardarNotas
				}
			}
		}
		fmt.Printf("%d\n", contarComposto)
	}

}
