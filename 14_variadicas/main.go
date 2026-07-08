package main

import "fmt"

func suma(numeros ...int) int {

	total := 0
	for _, num := range numeros {
		total += num
	}

	return total
}

func main() {
	println("Hola mundo")

	fmt.Println(suma(1, 2, 3, 4, 5, 6))

	arreglo := []int{6, 7, 8, 9, 10}
	fmt.Println(suma(arreglo...))

}
