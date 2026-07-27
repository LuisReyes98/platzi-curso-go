package main

import (
	"fmt"
	"time"
)

func main() {
	// For our example we'll select across two channels.

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 3; i++ {
		select {
		// Flujo el select llama al canal 1
		// el canal 1 llama a ejecutar su gorutine
		// luego que llega el mensaje del canal 1, se imprime en pantalla
		case msg1 := <-c1:
			fmt.Println("Canal 1:", msg1)
		case msg2 := <-c2:
			fmt.Println("Canal 2:", msg2)
		}
	}

}
