package main

import (
	"fmt"
)

func main() {

	musicNotes := [3]string{"do", "re", "mi"}
	nums := [5]float64{1.0, 2.0, 3.0, 4.0, 5.0}
	var sum float64

	for _, res := range nums {
		sum += res
	}

	media := sum / float64(len(nums))
	fmt.Printf("%.2f\n", media)

	for index, notes := range musicNotes {
		fmt.Printf("[%d] - %s\n", index, notes)
	}

	for _, notes := range musicNotes {
		fmt.Printf("%s\n", notes)
	}
}
