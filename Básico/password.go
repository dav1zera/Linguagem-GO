package main

import (
	"fmt"
)

const passwd = 2002

func main() {

	var n1 int
	fmt.Println("Enter your password: ")
	fmt.Scanf("%d", &n1)

	if n1 == passwd {
		fmt.Printf("Acess allow\n")
	} else {
		fmt.Printf("Acess Denied\n")
	}

}
