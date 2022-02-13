package main

import "fmt"

func main() {
	var i, j string = "Hello", "World"
	var n1, n2 int = 10, 20
	var nome string = "Davi"
	var idade int = 20

	//Print
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	fmt.Print(i, "\n", j, "\n")

	fmt.Print(n1, "\n", n2, "\n")
	fmt.Println(i, j)

	fmt.Printf("Seu nome eh: %v com tipo %T\n", nome, nome)
	fmt.Printf("Sua idade eh: %v com tipo %T", idade, idade)
}
