package main

import (
	"fmt"
)

func main() {
	var num, teste, count int64

	fmt.Scanf("%d", &num)

	for count = 1; count <= num; count++ {
		fmt.Scanf("%d", &teste)

		if teste == 0 {
			fmt.Printf("NULL\n")
		} else if teste <= 0 && teste%2 == 0 {
			fmt.Printf("EVEN NEGATIVE\n")
		} else if teste <= 0 && teste%2 == -1 {
			fmt.Printf("ODD NEGATIVE\n")
		} else if teste >= 0 && teste%2 == 0 {
			fmt.Printf("EVEN POSITIVE\n")
		} else if teste >= 0 && teste%2 == 1 {
			fmt.Printf("EVEN NEGATIVE\n")
		}

	}

	return

}
