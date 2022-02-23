package main

import (
	"fmt"
)

func main() {

	var creature string = "Shark"
	var pointer *string = &creature

	fmt.Println(creature)
	fmt.Println(pointer)
	fmt.Println(*pointer)

	*pointer = "Croc"
	fmt.Println(*pointer)

	fmt.Println(creature)

}
