package main

import (
	"fmt"
)

func main() {
	var humage int64
	fmt.Println("Digite sua idade: ")
	fmt.Scanf("%d", &humage)

	year := humage / 365
	months := (humage % 365) / 30
	days := (humage % 365) % 30

	fmt.Println(year)
	fmt.Println(months)
	fmt.Println(days)

}
