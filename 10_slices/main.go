package main

import (
	"fmt"
	"slices"
)

func main() {
	var nombres []string = []string{}

	fmt.Println("Nombres:", nombres, "Len:", len(nombres), "Es nulo:", nombres == nil)

	nombres = []string{"Zakamoto"}

	fmt.Println("Nombres:", nombres, "Len:", len(nombres), "Es nulo:", nombres == nil)

	// Transforma el arreglo en un slice
	nombres = make([]string, 3)

	nombres[0] = "Juan"
	nombres[1] = "Pedro"
	nombres[2] = "Maria"

	fmt.Println("Nombres:", nombres, "Len:", len(nombres), "Es nulo:", nombres == nil)
	fmt.Println("Segundo", nombres[1])

	nombres = append(nombres, "Luis", "Jose", "Ana")
	fmt.Println("Nombres:", nombres, "Len:", len(nombres), "Es nulo:", nombres == nil)

	arreglo2 := [3]string{"A", "B", "C"}

	arreglo3 := []string{"A", "B", "C"}

	if slices.Equal(arreglo2[:], arreglo3) {
		fmt.Println("Los arreglos son iguales")
	}
}
