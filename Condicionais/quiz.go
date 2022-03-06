package main

import (
	"fmt"
)

func main() {

	fmt.Printf("---------Bem vindo ao meu quiz---------\n")
	fmt.Printf("Quem venceu a copa de 1994?\n")
	fmt.Printf("A) Itália B) Brasil C) Alemanha D) Argentina\n")
	var escolha string
	fmt.Scanf("%s", &escolha)

	if escolha == "B" {
		fmt.Printf("Acertou\n")

	}

	if escolha == "A" {
		fmt.Printf("Errou\n")
	}

	if escolha == "C" {
		fmt.Printf("Errou\n")
	}

	if escolha == "D" {
		fmt.Printf("Errou\n")
	}

}
