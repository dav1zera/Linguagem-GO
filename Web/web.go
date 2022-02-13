package main

import (
	"fmt"
	"log"
	"net/http"
)

func request(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", request)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
