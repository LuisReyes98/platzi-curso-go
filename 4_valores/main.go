package main

import "fmt"

func main() {
	var numero int = 10
	fmt.Println(numero)
	numero2 := 20
	fmt.Println(numero2)

	numeroEntero := 30
	fmt.Println(numeroEntero)

	numeroFlotante := 3.14
	fmt.Println(numeroFlotante)

	resultado := float64(numeroEntero) + numeroFlotante
	fmt.Println(resultado)

	var nombre string = "Adam"
	fmt.Println(nombre)
	appellido := "Smith"
	fmt.Println(appellido)

	nombreCompleto := nombre + " " + appellido
	fmt.Println(nombreCompleto)
}
