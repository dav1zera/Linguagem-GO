package main

import (
	"fmt"
	"strings"
)

func main() {
	votos := "davi\nchibita\ncorró\nbalanegra\ngrampão\ncula"

	listavotos := strings.Split(votos, "\n")

	maps := make(map[string]int)

	for _, nomes := range listavotos {
		maps[nomes]++
	}

	fmt.Printf("%v \n", maps)

}
