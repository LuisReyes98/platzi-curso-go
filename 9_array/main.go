package main

import "fmt"

func main() {

	var lista [5]int

	fmt.Println("Lista:", lista)

	lista[4] = 100

	fmt.Println("Lista actualizada:", lista)

	fmt.Println("Tamaño de la lista:", len(lista))
}
