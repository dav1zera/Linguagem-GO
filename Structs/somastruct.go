package main

import (
	"fmt"
)

type Soma struct {
	n1 int
	n2 int
}

func main() {

	res := Soma{
		n1: 10,
		n2: 10,
	}

	fmt.Print(res.n1 + res.n2)

}
