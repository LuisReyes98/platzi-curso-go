package main

import (
	"fmt"
	"time"
)

func function(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Llamada a la función de manera síncrona

	go function("goroutine")
	function("direct")

	// si
	// You can also start a goroutine for an anonymous
	// function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	// Esperar a que la goroutine termine
	time.Sleep(time.Second)

}
