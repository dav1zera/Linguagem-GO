package main

import "fmt"

func main() {

	// arr1 := [5]int{} //not initialized
	// arr2 := [5]int{1,2} //partially initialized
	// arr3 := [5]int{1,2,3,4,5} //fully initialized
	// arr1 := [5]int{1:10,2:40}

	//Arrays e seus tamanhos
	var ints = [3]int{0, 1, 2}
	var nomes = [4]string{"Davi", "Alves", "Dos", "Santos"}
	var dinos = [...]string{"Rex", "Spino", "Oviraptor"}
	var times = [...]string{"Barcelona", "Real Madrid", "Ajax"}
	var jogadores = [...]string{"Messi", "Ronaldo", "Pelé", "Cruyff", "Maradona"}
	var cartas = [4]string{"Exodia", "Obelisco", "Slifer", "Rá"}

	fmt.Println(cartas)
	fmt.Println(jogadores)

	fmt.Println(len(ints))
	fmt.Println(len(nomes))
	fmt.Println(len(dinos))
	fmt.Println(ints)
	fmt.Println(nomes)
	fmt.Println(dinos)
	fmt.Println(times)

	fmt.Println(dinos[1])
	fmt.Println(nomes[2])
	fmt.Println(ints[2])

	dinos[1] = "Colossus"
	fmt.Println(dinos)

}
