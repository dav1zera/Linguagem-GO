package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func size(url string, c chan int) {
	fmt.Printf("Pegando o site: %s\n", url)
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("Tamanho: %d da pag %s\n", len(body), url)
	c <- len(body)
}

func main() {
	c := make(chan int)

	go size("https://www.tpbproxypirate.com/pt/", c)
	go size("https://www.sbt.com.br/", c)
	go size("https://www.grandepremio.com.br/f1/", c)
	go size("https://desciclopedia.org/wiki/Acre", c)
	var total int
	total += <-c
	total += <-c
	total += <-c
	total += <-c
	fmt.Printf("Tamanho total = %d\n", total)
	fmt.Printf("End\n")

}
