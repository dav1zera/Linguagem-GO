package main

import "fmt"

func main() {

	var n1, n2, res float64
	var menu int

	println("1-Soma 2-Sub 3-Div 4-Multi")
	fmt.Scan(&menu)

	switch menu {
	case 1:
		fmt.Println("Digite o primeiro número: ")
		fmt.Scan(&n1)
		fmt.Println("Digite o segundo número: ")
		fmt.Scan(&n2)
		res = n1 + n2
		fmt.Printf("A soma é: %.2f ", res)
		break

	case 2:
		fmt.Println("Digite o primeiro número: ")
		fmt.Scan(&n1)
		fmt.Println("Digite o segundo número: ")
		fmt.Scan(&n2)
		res = n1 - n2
		fmt.Printf("A sub é: %.2f ", res)
		break

	case 3:
		fmt.Println("Digite o primeiro número: ")
		fmt.Scan(&n1)
		fmt.Println("Digite o segundo número: ")
		fmt.Scan(&n2)
		res = n1 / n2
		fmt.Printf("A div é: %.2f ", res)
		break

	case 4:
		fmt.Println("Digite o primeiro número: ")
		fmt.Scan(&n1)
		fmt.Println("Digite o segundo número: ")
		fmt.Scan(&n2)
		res = n1 * n2
		fmt.Printf("A multi é: %.2f ", res)
		break

	default:
		println("Escolha incorreta")

	}

}
