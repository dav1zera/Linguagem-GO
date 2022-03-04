package main

import (
	"fmt"
	"log"
)

var rendimentoTinta float64

func ola() {
	fmt.Println("Teste")

}

func calcTinta(larg, alt float64) (float64, error) {
	if larg <= 0 || alt <= 0 {
		return 0, fmt.Errorf("A largura e altura devem ser positivos")
	}

	area := larg * alt
	return area / rendimentoTinta, nil

}

func main() {
	ola()
	rendimentoTinta = 12.1
	resultado, err := calcTinta(4.2, 11.1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%10.4v\n", resultado)

}
