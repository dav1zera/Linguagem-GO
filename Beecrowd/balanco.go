package main

import (
	"fmt"
)

func main() {

	var expressão, caractere string
	var qtdParentesesAbertos, contador int64

	for {
		_, errString := fmt.Scanf("%s\n", &expressão)
		if errString != nil {
			break
		}

		qtdParentesesAbertos = 0

		for contador = 0; contador < int64(len(expressão)); contador++ {
			caractere = string(expressão[contador])

			if caractere == "(" {
				qtdParentesesAbertos++
			} else if caractere == ")" {
				qtdParentesesAbertos--

				if qtdParentesesAbertos < 0 {
					break
				}
			}

		}

		if qtdParentesesAbertos == 0 {
			fmt.Printf("correct\n")
		} else {
			fmt.Printf("incorret\n")
		}

	}

}
