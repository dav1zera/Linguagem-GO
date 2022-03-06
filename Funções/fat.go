package main

import (
	"fmt"
	"log"
)

func calcFat(num int64) (int64, error) {
	var fat int64
	if num < 0 {
		return 0, fmt.Errorf("Deve ser positivo")
	}

	for fat = 1; num > 1; num = num - 1 {
		fat = fat * num
	}

	return fat, nil
}

func main() {

	var num1 int64
	fmt.Scanf("%d", &num1)
	res, err := calcFat(num1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("O fat é: %d", res)

}
