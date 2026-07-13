package main

import "fmt"

func modificarArreglo(arreglo *[]int) {
	(*arreglo)[0] = 100
}

func main() {
	datos := []int{1, 2, 3, 4, 5}

	fmt.Println("numeros: ", datos)

	modificarArreglo(&datos)
	fmt.Println("numeros: ", datos)

}
