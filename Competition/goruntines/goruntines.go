package main

import (
	"fmt"
	"time"
)

func main() {

	go hello("Msg")
	hello("Oi")

}

func hello(msg string) {

	for {
		fmt.Println(msg)
		time.Sleep(time.Second)
	}

}
