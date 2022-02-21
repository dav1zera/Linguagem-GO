package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	leitura := bufio.NewReader(os.Stdin)
	fmt.Println("Entre com uma String: ")
	str, _ := leitura.ReadString('\n')
	tamanho := 0

	for _, i := range str {
		fmt.Printf("%c", i)
		tamanho++
	}

	fmt.Println("\no tamanho da string é: ", tamanho-1)

}
