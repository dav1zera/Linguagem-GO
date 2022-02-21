package main

import (
	"fmt"
)

type Animal struct {
	name string
	tipo string
}

func main() {

	animal1 := Animal{
		name: "Cobra",
		tipo: "Carnivoro",
	}

	fmt.Println(animal1.name, animal1.tipo)

}
