package main

import (
	"fmt"
	"main/data"
)

type event struct {
	title string
	data.Date
}

func main() {

	formatura := event{title: "Dia da formatura"}
	formatura.SetDay(11)
	formatura.SetMonth(02)
	formatura.SetYear(2024)
	fmt.Printf("%v \n", formatura)

}
