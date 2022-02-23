package main

//Ponteiros

import (
	"fmt"
)

func main() {

	a := 10

	var pointer *int = &a
	//fmt.Println(pointer)
	//fmt.Println(*pointer)
	//fmt.Println(&pointer)

	*pointer = 50
	b := &a
	*b = 60
	fmt.Println(a)

}
