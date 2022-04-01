package main

import (
	"fmt"
)

func main() {

	var dates string
	fmt.Scanf("%s", &dates)

	DD := dates[0:2]
	MM := dates[3:5]
	AA := dates[6:8]

	fmt.Printf("%s/%s/%s\n", MM, DD, AA)
	fmt.Printf("%s/%s/%s\n", AA, MM, DD)
	fmt.Printf("%s-%s-%s\n", DD, MM, AA)

}
