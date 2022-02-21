package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Digite a primeira nota: ")
	readernota1 := bufio.NewReader(os.Stdin)
	inputnota1, errnota1 := readernota1.ReadString('\n')
	fmt.Println("Digite a Segunda nota: ")
	readernota2 := bufio.NewReader(os.Stdin)
	inputnota2, errnota2 := readernota2.ReadString('\n')

	if errnota1 != nil {
		log.Fatal(errnota1)
	}

	if errnota2 != nil {
		log.Fatal(errnota2)
	}

	inputnota1 = strings.TrimSpace(inputnota1)
	nota1, _ := strconv.ParseFloat(inputnota1, 64)
	inputnota2 = strings.TrimSpace(inputnota2)
	nota2, _ := strconv.ParseFloat(inputnota2, 64)

	media := (nota1 + nota2) / 2

	var status string
	if media >= 7 {
		status = "Aprovado"
	} else if media < 3 {
		status = "Reprovado"
	} else {
		status = "Final"
	}
	fmt.Println("Sua nota foi:", media, "\nSituação:", status)

}
