package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const saludoEnTailandes = "สวัสดีครับ"

	contarRunas(saludoEnTailandes)
	contarRunas("Hola")
}

func contarRunas(saludo string) {
	fmt.Println("El saludo en tailandés es:", saludo)

	fmt.Println("Len numero de bytes:", len(saludo))
	fmt.Println("Runas numero de caracteres:", utf8.RuneCountInString(saludo))

	for i := 0; i < len(saludo); i++ {
		fmt.Printf("%x ", saludo[i])
	}

	for idx, valorRuna := range saludo {
		fmt.Println("%#U comienza en %d\n", valorRuna, idx)
	}

}
