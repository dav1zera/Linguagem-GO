package main

import (
	"fmt"
)

func isDDD(city int64) {

	if city == 61 {
		fmt.Println("São Paulo")
	} else if city == 71 {
		fmt.Println("Salvador")
	} else if city == 21 {
		fmt.Println("Rio de Janeiro")
	} else {
		fmt.Println("Erro")
	}

}

func main() {
	var ddd int64
	fmt.Scanf("%i", &ddd)

	isDDD(ddd)
	fmt.Println(ddd)

}
