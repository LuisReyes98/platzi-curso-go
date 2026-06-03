package main

// Condicionales if y else en Go con operadores básicos
import "fmt"

func main() {
	nombre := "Juan"
	edad := 25

	if edad >= 18 {
		fmt.Printf("%s es mayor de edad.\n", nombre)
	} else {
		fmt.Printf("%s es menor de edad.\n", nombre)
	}

	if 8%2 == 0 {
		fmt.Println("8 es un número par.")
	}

	if numero, x := edad/3, 0; numero < 5 && x == 0 {
		fmt.Printf("%d es menor que 5.\n", numero)
	} else if numero < 10 {
		fmt.Printf("%d es menor que 10 pero mayor o igual a 5.\n", numero)
	}
}
