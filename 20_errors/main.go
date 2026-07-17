package main

import (
	"errors"
	"fmt"
)

var ErrorMio = fmt.Errorf("errore mio")
var ErrorMioDeCafe = errors.New("error de cafe")

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No error")
	}
}

func hacerAlgo(args int) error {

	switch args {
	case 2:
		return ErrorMio

	case 3:
		return ErrorMioDeCafe
	}
	return fmt.Errorf("error con formato: %d", args)
}

func main() {
	// checkError(ErrorMio)
	// checkError(ErrorMioDeCafe)

	for i := 0; i < 5; i++ {
		err := hacerAlgo(i)
		if err != nil {
			// panic(err)
			// panic("algo salio mal")
		}
		if errors.Is(err, ErrorMioDeCafe) {
			fmt.Println("El error del cafe ocurrio")
		}
		checkError(err)
	}
}
