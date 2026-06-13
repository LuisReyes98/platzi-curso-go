package main

import (
	"fmt"
	"maps"
)

func main() {
	myMap := make(map[string]int)

	myMap["Juan"] = 25
	myMap["Jose"] = 30
	myMap["Carmen"] = 32
	fmt.Println("Mi mapa:", myMap)

	juan := myMap["Juan"]
	fmt.Println("Juan:", juan)

	// Length
	fmt.Println("Len", len(myMap))

	valor, exists := myMap["Juan"]
	fmt.Println("My mapa juan:", valor, exists)

	delete(myMap, "Juan")
	fmt.Println("Mi mapa:", myMap)

	clear(myMap)
	fmt.Println("Mi mapa:", myMap)

	map1 := map[string]int{
		"hola": 1,
	}
	map2 := map[string]int{
		"hola": 1,
	}

	if maps.Equal(map1, map2) {
		fmt.Println("Los mapas son iguales")
	}

}
