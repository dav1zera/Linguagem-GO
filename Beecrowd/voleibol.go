package main

import (
	"fmt"
)

func main() {

	var n, s, b, a, s1, b1, a1, i float64
	var sum_s float64 = 0.00
	var sum_b float64 = 0.00
	var sum_a float64 = 0.00
	var sum_s1 float64 = 0.00
	var sum_b1 float64 = 0.00
	var sum_a1 float64 = 0.00
	var ans_s, ans_b, ans_a float64
	var name string

	fmt.Scanf("%d", &n)

	for i = 1; i <= n; i++ {
		fmt.Scanf("%s", &name)
		fmt.Scanf("%f %f %f", &s, &b, &a)
		fmt.Scanf("%f %f %f", &s1, &b1, &a1)

		sum_s += s
		sum_b += b
		sum_a += a
		sum_s1 += s1
		sum_b1 += b1
		sum_a1 += a1
	}

	ans_s = (sum_s1 / sum_s) * 100.00
	ans_b = (sum_b1 / sum_b) * 100.00
	ans_a = (sum_a1 / sum_a) * 100.00

	fmt.Printf("Pontos de Saque: %.2f%%.\n", ans_s)
	fmt.Printf("Pontos de Bloqueio: %.2f%%.\n", ans_b)
	fmt.Printf("Pontos de Ataque: %.2f%%.\n", ans_a)
}
