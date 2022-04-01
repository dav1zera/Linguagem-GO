package main

import (
	"fmt"
	"time"
)

//Concerrencia e Paralelismo
//chan = channel
//Sprintf = criar uma string
func simple(msg string, c chan string) {
	for i := 0; i < 5; i++ {

		c <- fmt.Sprintf(msg)

	}
}

func main() {
	c := make(chan string)
	go simple("Ping", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("Recebido: %q \n", <-c)
		fmt.Printf("Pong\n")
		time.Sleep(1 * time.Second)
	}

	time.Sleep(10 * time.Second)
	fmt.Println("End")
}
