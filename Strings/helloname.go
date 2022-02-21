package main

import (
	"fmt"
)

func main() {
	var nome string
	fmt.Println("Digite seu nome: ")
	fmt.Scan(&nome)
	fmt.Printf("Olá %s", nome)
}
