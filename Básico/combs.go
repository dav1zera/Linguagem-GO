//Básico com tratamento de erros
//calc := tempoGasto * velocidadeMedia / 12.0;

package main

import (
	"fmt"
	"log"
)

const kmL = 12

func calcCombust(tempoGasto, velocidadeMedia float64) (float64, error) {

	if tempoGasto <= 0 || velocidadeMedia <= 0 {
		return 0, fmt.Errorf("Devem ser positivos")
	}

	calcKm := tempoGasto * velocidadeMedia / kmL
	return calcKm, nil

}

func main() {

	var n1, n2 float64
	fmt.Println("Digite um número: ")
	fmt.Scanf("%f", &n1)
	fmt.Println("Digite o segundo número: ")
	fmt.Scanf("%f", &n2)

	gastoCarro, err := calcCombust(n1, n2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("O gasto foi %.3f", gastoCarro)

}
