package main

import (
	"fmt"
)

func abc(a *int) {
	*a = 200

}

func main() {
	variavel := 10

	abc(&variavel)
	fmt.Println(variavel)
}
