package main

import (
	"fmt"
)

type Calculadora struct {
	n1 float64
	n2 float64
}

func (res *Calculadora) Soma() {
	fmt.Println("A soma é: ", res.n1+res.n2)
}

func (res *Calculadora) Sub() {
	fmt.Println("A subtração é: ", res.n1-res.n2)
}

func (res *Calculadora) Multi() {
	fmt.Println("A multiplicação é: ", res.n1*res.n2)
}

func (res *Calculadora) Div() {
	fmt.Println("A divisão é: ", res.n1/res.n2)
}

func main() {
	var num1, num2 float64
	fmt.Println("Digite o primeiro número: ")
	fmt.Scanf("%d", &num1)
	fmt.Println("Digite o segundo: ")
	fmt.Scanf("%d", &num2)

	res = Calculadora{
		n1: num1,
		n2: num2,
	}

	op := 1

	for op >= 1 {
		fmt.Println("1 - Adição: ")
		fmt.Println("2 - Multiplicação: ")
		fmt.Println("3 - Divisão: ")
		fmt.Println("4 - Subtração: ")
		fmt.Print("5 - Sair: ")
		fmt.Scanf("%d", &op)

		switch op {
		case 1:
			res.Soma()
		case 2:
			res.Multi()
		case 3:
			res.Div()
		case 4:
			res.Sub()
		case 5:
			op = 0
			break
		default:
			fmt.Println("Escolha errada")

		}
	}

}
