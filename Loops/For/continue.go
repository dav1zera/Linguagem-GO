package main

import (
	"fmt"
)

func main() {
	for ini := 0; ini < 1100; ini++ {
		if ini != 0 && ini%3 == 0 && ini%7 == 0 && ini%9 == 0 {
			fmt.Println(ini)
			break
		}
	}
}
