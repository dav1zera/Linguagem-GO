package main

import (
	"fmt"
	"time"
)

func main() {

	var dates [5]time.Time
	dates[0] = time.Now()
	dates[1] = time.Now()
	dates[2] = time.Unix(120, 0)
	dates[3] = time.Now()
	dates[4] = time.Now()

	fmt.Println(dates[1])
}
