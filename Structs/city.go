package main

import (
	"fmt"
)

type Cidades struct {
	adress     string
	postalcode int64
}

func main() {

	rua := Cidades{
		adress:     "Rua Gilberto Valentin",
		postalcode: 123,
	}

	fmt.Println(rua)
}
