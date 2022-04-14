package main

import (
	"fmt"
)

func main() {

	var i int64
	var M int64
	var N int64

	var soma int64 = 0
	var contador int64 = 1

	for contador != 0 {

		fmt.Scanf("%d %d", &M, &N)

		if M > N {
			aux := M
			M = N
			N = aux

		}

		if M <= 0 || N <= 0 {
			contador--
		}

		if contador != 0 {
			soma = 0
			for i = M; i <= N; i++ {

				fmt.Printf("%d ", i)
				soma = soma + i
			}
			fmt.Printf("Sum=%d\n", soma)
		}
	}

}
