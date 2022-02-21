package main

import "fmt"

func main() {

	var n1, n2, n3, n4, dif int64
	fmt.Scanf("%d %d %d %d", &n1, &n2, &n3, &n4)

	dif = (n1*n2 - n3*n4)
	fmt.Printf("DIF = %d", dif)

}
