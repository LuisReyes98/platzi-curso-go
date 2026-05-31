package main

// No puedes compilar una app
// si no usas todos los paquetes importados
import (
	"fmt"
	"os"
)

func main() {
	envVar := os.Getenv("GO_ENV")

	if envVar == "" {
		fmt.Println("La variable de entorno GO_ENV no está establecida.")
	} else {
		fmt.Printf("La variable de entorno GO_ENV está establecida en: %s\n", envVar)
	}

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Printf("Error al crear el archivo: %v\n", err)
		return
	}
	defer file.Close()
}
