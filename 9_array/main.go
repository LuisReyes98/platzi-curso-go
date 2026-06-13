package main

import "fmt"

func main() {

	var lista [5]int

	fmt.Println("Lista:", lista)

	lista[4] = 100

	fmt.Println("Lista actualizada:", lista)

	fmt.Println("Tamaño de la lista:", len(lista))

	myLista := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Mi lista:", myLista)

	listaInferida := [...]int{1, 4, 5, 6, 7}
	fmt.Println("Lista sin límite:", listaInferida)
	// Length
	fmt.Println("Len", len(listaInferida))

	limitlessList := []int{}
	// limitlessList = []int{}
	fmt.Println("Lista sin límite:", limitlessList)
	limitlessList = append(limitlessList, 4)
	limitlessList = append(limitlessList, 3)
	limitlessList = append(limitlessList, 5)
	fmt.Println("Lista sin límite:", limitlessList)
	fmt.Println("Len", len(limitlessList))
}
