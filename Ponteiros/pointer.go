package main

import (
	"fmt"
)

func main() {

	var creature string = "Shark"
	var pointer *string = &creature

	fmt.Println("", creature)
	fmt.Println("", pointer)

}
