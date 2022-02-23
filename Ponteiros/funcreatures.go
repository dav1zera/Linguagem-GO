package main

import (
	"fmt"
)

type Creature struct {
	Species string
}

func main() {
	var creature Creature = Creature{
		Species: "Sharks",
	}

	fmt.Printf("%+v\n", creature)
	changeCreature(creature)
	fmt.Printf("%+v\n", creature)
}

func changeCreature(creature Creature) {
	creature.Species = "Crocs"
	fmt.Printf("%+v\n", creature)

}
